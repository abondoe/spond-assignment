import {
  Select,
  SelectContent,
  SelectItem,
  SelectTrigger,
  SelectValue,
} from "./ui/select";
import { type MemberType } from "../generated-types/api";
import { Field, FieldDescription, FieldLabel } from "./ui/field";

interface MemberTypeSelectProps {
  memberTypes: MemberType[] | undefined;
  value: string | null;
  onChange: (value: string | null) => void;
  error: string;
}

function MemberTypeSelect({
  memberTypes,
  value,
  onChange,
  error,
}: MemberTypeSelectProps) {
  // Finn objektet basert på propen 'value' i stedet for intern state
  const selectedMemberType = memberTypes?.find(
    (t) => t.id.toString() === value,
  );

  return (
    <div className="space-y-4">
      <Field>
        <FieldLabel htmlFor="select-member-type">Member type</FieldLabel>
        <Select value={value} onValueChange={onChange}>
          <SelectTrigger
            id="select-member-type"
            className="w-full max-w-48"
            aria-invalid={!!error}
          >
            <SelectValue>
              {selectedMemberType
                ? selectedMemberType.name
                : "Select member type"}
            </SelectValue>
          </SelectTrigger>
          <SelectContent>
            {memberTypes?.map((type) => (
              <SelectItem key={type.id} value={type.id.toString()}>
                {type.name}
              </SelectItem>
            ))}
          </SelectContent>
        </Select>
        {error && (
          <FieldDescription className="text-destructive">
            {error}
          </FieldDescription>
        )}
      </Field>
    </div>
  );
}

export default MemberTypeSelect;
