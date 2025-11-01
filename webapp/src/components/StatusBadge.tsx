export default function StatusBadge({ status }: { status: string }) {
  const colors: Record<string, string> = {
    active: "bg-green-100 text-green-700",
    blocked: "bg-red-100 text-red-700",
    on_leave: "bg-yellow-100 text-yellow-700",
    terminated: "bg-gray-200 text-gray-600",
    rehired: "bg-blue-100 text-blue-700",
  };

  const label = {
    active: "Ativo",
    blocked: "Bloqueado",
    on_leave: "Licença",
    terminated: "Demitido",
    rehired: "Recontratado",
  }[status] || status;

  return (
    <span className={`px-3 py-1 text-xs font-medium rounded-full ${colors[status] || "bg-gray-100 text-gray-500"}`}>
      {label}
    </span>
  );
}
