import { ConnectionService } from '@/neosync-api-client/mgmt/v1alpha1/connection_connect';
import { JobService } from '@/neosync-api-client/mgmt/v1alpha1/job_connect';
import { UserAccountService } from '@/neosync-api-client/mgmt/v1alpha1/user_account_connect';
import {
  Code,
  ConnectError,
  Interceptor,
  PromiseClient,
  Transport,
  createPromiseClient,
} from '@connectrpc/connect';
import { createConnectTransport } from '@connectrpc/connect-node';
import { getToken } from 'next-auth/jwt';
import { NextRequest, NextResponse } from 'next/server';

interface NeosyncContext {
  connectionClient: PromiseClient<typeof ConnectionService>;
  userClient: PromiseClient<typeof UserAccountService>;
  jobsClient: PromiseClient<typeof JobService>;
}

type NeosyncApiHandler<T = unknown> = (ctx: NeosyncContext) => Promise<T>;

interface ErrorMessageResponse {
  message: string;
}

export function withNeosyncContext<T = unknown>(
  handler: NeosyncApiHandler<T>
): (req: NextRequest) => Promise<NextResponse<T | ErrorMessageResponse>> {
  return async (req) => {
    try {
      const output = await handler(await getNeosyncContext(req));
      return NextResponse.json(output);
    } catch (err) {
      if (err instanceof ConnectError) {
        return NextResponse.json(
          { message: err.message },
          {
            status: translateGrpcCodeToHttpCode(err.code),
          }
        );
      }
      return NextResponse.json(
        {
          message: 'unknown error type',
        },
        {
          status: 500,
        }
      );
    }
  };
}

async function getNeosyncContext(req: NextRequest): Promise<NeosyncContext> {
  // const res = new NextResponse();

  const jwt = await getToken({ req: req, raw: true });
  if (!jwt) {
    throw new Error('no session provided');
  }

  const transport = getAuthenticatedConnectTransport(
    getApiBaseUrlFromEnv(),
    () => jwt
  );

  return {
    connectionClient: createPromiseClient(ConnectionService, transport),
    userClient: createPromiseClient(UserAccountService, transport),
    jobsClient: createPromiseClient(JobService, transport),
  };
}

function getAuthenticatedConnectTransport(
  baseUrl: string,
  getAccessToken: () => Promise<string> | string
): Transport {
  return createConnectTransport({
    baseUrl,
    httpVersion: '2',
    interceptors: [getAuthInterceptor(getAccessToken)],
  });
}

function getAuthInterceptor(
  getAccessToken: () => Promise<string> | string
): Interceptor {
  return (next) => async (req) => {
    const accessToken = await getAccessToken();
    req.header.set('Authorization', `Bearer ${accessToken}`);
    return next(req);
  };
}

function getApiBaseUrlFromEnv(): string {
  const apiUrl = process.env.NEOSYNC_API_BASE_URL;
  if (!apiUrl) {
    throw new Error('must provide NEOSYNC_API_BASE_URL');
  }
  return apiUrl;
}

function translateGrpcCodeToHttpCode(code: Code): number {
  switch (code) {
    case Code.InvalidArgument:
    case Code.FailedPrecondition:
    case Code.OutOfRange: {
      return 400;
    }
    case Code.Unauthenticated: {
      return 401;
    }
    case Code.PermissionDenied: {
      return 403;
    }
    case Code.Unimplemented:
    case Code.NotFound: {
      return 404;
    }
    case Code.AlreadyExists: {
      return 409;
    }
    case Code.Unavailable: {
      return 503;
    }
    default: {
      return 500;
    }
  }
}
