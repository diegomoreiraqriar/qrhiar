// src/utils/auditUtils.ts
export const getActionLabel = (action: string): string => {
  const labels: Record<string, string> = {
    activate: "Usuário Ativado",
    block: "Usuário Bloqueado",
    leave: "Usuário em Licença",
    promote: "Usuário Promovido",
    terminate: "Usuário Desligado",
    joiner: "Usuário Criado",
    update: "Perfil Atualizado",
  };
  return labels[action] || action;
};

export const getActionColor = (action: string): string => {
  const colors: Record<string, string> = {
    activate: "bg-green-100 text-green-800",
    block: "bg-red-100 text-red-800",
    leave: "bg-yellow-100 text-yellow-800",
    promote: "bg-blue-100 text-blue-800",
    terminate: "bg-gray-200 text-gray-700",
    joiner: "bg-purple-100 text-purple-800",
    update: "bg-indigo-100 text-indigo-800",
  };
  return colors[action] || "bg-gray-100 text-gray-700";
};
