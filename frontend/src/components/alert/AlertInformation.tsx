import { InfoIcon } from "lucide-react";

import { Alert, AlertDescription, AlertTitle } from "../ui/alert";

interface AlertInformationProps {
  title: React.ReactNode;
  description: React.ReactNode;
}

export function AlertInformation({
  title,
  description,
}: AlertInformationProps) {
  return (
    <Alert className="max-w-md">
      <InfoIcon />
      <AlertTitle>{title}</AlertTitle>
      <AlertDescription>{description}</AlertDescription>
    </Alert>
  );
}
