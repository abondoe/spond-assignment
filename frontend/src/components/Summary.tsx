import type { UserData } from "./UserInfoForm";

interface SummaryProps {
  memberType: string | undefined;
  userData: UserData;
}

function Summary({ memberType, userData }: SummaryProps) {
  const birthdate = new Date(userData.birthDate);

  return (
    <div className="space-y-6">
      <div className="bg-slate-50 p-4 rounded-lg border border-slate-100">
        <dl className="grid grid-cols-1 sm:grid-cols-2 gap-y-4 gap-x-8">
          <div className="flex flex-col">
            <dt className="text-xs text-slate-500 uppercase tracking-wider">
              Member type
            </dt>
            <dd className="text-sm font-semibold">{memberType}</dd>
          </div>

          <div className="flex flex-col">
            <dt className="text-xs text-slate-500 uppercase tracking-wider">
              Name
            </dt>
            <dd className="text-sm font-semibold">{userData.name}</dd>
          </div>

          <div className="flex flex-col">
            <dt className="text-xs text-slate-500 uppercase tracking-wider">
              Email
            </dt>
            <dd className="text-sm font-semibold truncate">{userData.email}</dd>
          </div>

          <div className="flex flex-col">
            <dt className="text-xs text-slate-500 uppercase tracking-wider">
              Phone number
            </dt>
            <dd className="text-sm font-semibold">{userData.phoneNumber}</dd>
          </div>

          <div className="flex flex-col">
            <dt className="text-xs text-slate-500 uppercase tracking-wider">
              Birth date
            </dt>
            <dd className="text-sm font-semibold">
              {birthdate.toLocaleDateString()}
            </dd>
          </div>
        </dl>
      </div>
    </div>
  );
}

export default Summary;
