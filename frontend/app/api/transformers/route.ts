import { withNeosyncContext } from '@/api-only/neosync-context';
import { GetTransformersRequest } from '@/neosync-api-client/mgmt/v1alpha1/job_pb';
import { NextRequest, NextResponse } from 'next/server';

export async function GET(req: NextRequest): Promise<NextResponse> {
  return withNeosyncContext(async (ctx) => {
    return ctx.jobsClient.getTransformers(new GetTransformersRequest({}));
  })(req);
}