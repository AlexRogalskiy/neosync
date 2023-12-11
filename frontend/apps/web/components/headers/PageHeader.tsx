import { cn } from '@/libs/utils';
import { ReactNode } from 'react';
import { Badge } from '../ui/badge';
import { Separator } from '../ui/separator';

interface Props {
  header: string;
  description?: string;
  leftBadgeValue?: string;
  extraHeading?: ReactNode;
  leftIcon?: ReactNode;
  progressSteps?: JSX.Element;
  pageHeaderContainerClassName?: string;
}

export default function PageHeader(props: Props) {
  const {
    header,
    description,
    extraHeading,
    leftIcon,
    pageHeaderContainerClassName,
    leftBadgeValue,
    progressSteps,
  } = props;
  return (
    <div
      className={cn(
        'page-header-container flex flex-col gap-2',
        pageHeaderContainerClassName
      )}
    >
      <div className="flex flex-row items-center justify-between">
        <div className="flex flex-row items-center gap-3">
          {leftIcon ? leftIcon : null}
          <div className="flex flex-row items-center gap-4">
            <h1 className="text-2xl font-bold tracking-tight">{header}</h1>
            {leftBadgeValue && (
              <Badge variant="outline">{leftBadgeValue}</Badge>
            )}
          </div>
        </div>

        <div className="flex-1 flex justify-center">{progressSteps}</div>

        {extraHeading ? <div>{extraHeading}</div> : null}
      </div>

      {description ? (
        <h3 className="text-muted-foreground text-sm">{description}</h3>
      ) : null}

      <Separator />
    </div>
  );
}