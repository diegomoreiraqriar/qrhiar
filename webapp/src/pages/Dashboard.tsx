import { useEffect, useState } from "react";
import { getUsers } from "../api/users";
import { getActionLabel, getActionColor } from "../utils/auditUtils";
import Navbar from "../components/Navbar";
import toast from "react-hot-toast";
import CreateUserModal from "../components/CreateUserModal";
import api from "../api/client";

interface Manager {
  displayName: string;
}

interface Company {
  name: string;
}

interface AuditLog {
  id: string;
  action: string;
  reason: string;
  old_value: string;
  new_value: string;
  created_at: string;
}

interface User {
  id: string;
  name: string;
  email: string;
  position: string;
  status: string;
  company?: Company;
  manager?: Manager;
}

export default function Dashboard() {
  const [users, setUsers] = useState<User[]>([]);
  const [loading, setLoading] = useState(true);
  const [selectedUser, setSelectedUser] = useState<User | null>(null);
  const [logs, setLogs] = useState<AuditLog[]>([]);
  const [loadingLogs, setLoadingLogs] = useState(false);
  const [updating, setUpdating] = useState(false);
  const [showCreateModal, setShowCreateModal] = useState(false);

  // 🔄 Buscar usuários
  const loadUsers = async () => {
    setLoading(true);
    try {
      const data = await getUsers();
      setUsers(data || []);
    } catch (err) {
      console.error("Erro ao buscar usuários:", err);
      toast.error("Erro ao carregar usuários");
    } finally {
      setLoading(false);
    }
  };

  useEffect(() => {
    loadUsers();
  }, []);

  // 🎨 Cor do card conforme status
  const getCardColor = (status: string): string => {
    switch (status) {
      case "active":
        return "bg-green-50 border-l-4 border-green-400 hover:bg-green-100";
      case "blocked":
        return "bg-red-50 border-l-4 border-red-400 hover:bg-red-100";
      case "inactive":
        return "bg-gray-50 border-l-4 border-gray-300 hover:bg-gray-100";
      case "leave":
        return "bg-yellow-50 border-l-4 border-yellow-400 hover:bg-yellow-100";
      default:
        return "bg-white border-l-4 border-gray-100";
    }
  };

  // 📜 Buscar logs
  const fetchUserLogs = async (userId: string) => {
    setLoadingLogs(true);
    try {
      const res = await api.get(`/api/third-parties/${userId}/logs`);
      const data = res.data;
      setLogs(data.logs || []);
    } catch (err) {
      console.error("Erro ao buscar logs:", err);
      toast.error("Erro ao carregar histórico");
    } finally {
      setLoadingLogs(false);
    }
  };

  //  Ações de status
  const handleAction = async (userId: string, action: string) => {
    setUpdating(true);

    const actionNames: Record<string, string> = {
      block: "Bloquear usuário",
      activate: "Ativar usuário",
      leave: "Colocar em férias",
      inactive: "Demitir usuário",
    };

    try {
      // ✅ Envia requisição PATCH autenticada com token (via api interceptor)
      const res = await api.patch(`/api/third-parties/${userId}/status`, { action });
      const data = res.data;

      // Atualiza lista principal
      setUsers((prev) =>
        prev.map((u) => (u.id === userId ? { ...u, status: data.user.status } : u))
      );

      // Atualiza usuário do modal (se aberto)
      if (selectedUser?.id === userId) {
        setSelectedUser({ ...selectedUser, status: data.user.status });
        fetchUserLogs(userId);
      }

      toast.success(`${actionNames[action]} realizada com sucesso!`);
    } catch (err) {
      console.error("Erro ao executar ação:", err);
      toast.error("Falha ao executar ação. Verifique os logs.");
    } finally {
      setUpdating(false);
    }
  };


  const handleViewDetails = (user: User) => {
    setSelectedUser(user);
    fetchUserLogs(user.id);
  };

  const closeModal = () => {
    setSelectedUser(null);
    setLogs([]);
  };

  return (
    <div className="min-h-screen bg-gray-50 text-gray-800">
      <Navbar />

      <main className="p-8">
        {/* Header */}
        <div className="flex justify-between items-center mb-6">
          <h1 className="text-2xl font-semibold text-gray-800">Colaboradores</h1>
          <button
            onClick={() => setShowCreateModal(true)}
            className="px-4 py-2 bg-indigo-600 text-white rounded-lg shadow hover:bg-indigo-700"
          >
            + Novo Usuário
          </button>
        </div>

        {/* Lista de usuários */}
        {loading ? (
          <p className="text-gray-500 text-center mt-12">Carregando usuários...</p>
        ) : users.length === 0 ? (
          <p className="text-gray-500 text-center mt-12">Nenhum usuário encontrado.</p>
        ) : (
          <div className="grid grid-cols-1 md:grid-cols-2 lg:grid-cols-3 gap-6">
            {users.map((user) => (
              <div
                key={user.id}
                className={`shadow-md transition rounded-xl p-6 border ${getCardColor(
                  user.status
                )}`}
              >
                <div className="flex items-center justify-between">
                  <h2 className="text-lg font-semibold text-gray-900">{user.name}</h2>
                  <span
                    className={`px-2 py-1 text-xs rounded-full ${
                      user.status === "active"
                        ? "bg-green-100 text-green-700"
                        : user.status === "inactive"
                        ? "bg-gray-200 text-gray-600"
                        : user.status === "blocked"
                        ? "bg-red-100 text-red-700"
                        : "bg-yellow-100 text-yellow-700"
                    }`}
                  >
                    {user.status === "active"
                      ? "Ativo"
                      : user.status === "inactive"
                      ? "Inativo"
                      : user.status === "blocked"
                      ? "Bloqueado"
                      : "Licença"}
                  </span>
                </div>

                <p className="text-sm text-gray-600 mt-1">{user.position}</p>

                <div className="mt-4 space-y-2 text-sm text-gray-700">
                  {user.email && <p>📧 <span className="font-medium">{user.email}</span></p>}
                  {user.company?.name && <p>🏢 <span className="font-medium">{user.company.name}</span></p>}
                  {user.manager?.displayName && <p>👤 <span className="font-medium">{user.manager.displayName}</span></p>}
                </div>

                <button
                  onClick={() => handleViewDetails(user)}
                  className="mt-4 w-full text-sm font-medium text-indigo-600 hover:text-indigo-800"
                >
                  Ver Detalhes
                </button>
              </div>
            ))}
          </div>
        )}
      </main>

      {/* Modal de criação */}
      {showCreateModal && (
        <CreateUserModal
          onClose={() => setShowCreateModal(false)}
          onUserCreated={loadUsers}
        />
      )}

      {/* Modal de detalhes */}
      {selectedUser && (
        <div className="fixed inset-0 bg-black bg-opacity-50 flex items-center justify-center p-4 z-50">
          <div className="bg-white rounded-xl w-full max-w-2xl shadow-lg overflow-hidden">
            <div className="p-6 border-b border-gray-200 flex justify-between items-center">
              <h2 className="text-xl font-semibold text-gray-900">
                {selectedUser.name}
              </h2>
              <button
                onClick={closeModal}
                className="text-gray-400 hover:text-gray-700 text-2xl"
              >
                ✖
              </button>
            </div>

            <div className="p-6 space-y-4">
              <p><strong>Email:</strong> {selectedUser.email || "—"}</p>
              <p><strong>Cargo:</strong> {selectedUser.position}</p>
              <p><strong>Status:</strong> {selectedUser.status}</p>
              <p><strong>Empresa:</strong> {selectedUser.company?.name}</p>
              <p><strong>Gestor:</strong> {selectedUser.manager?.displayName}</p>

              {/* Botões de ação */}
              <div className="flex flex-wrap gap-2 mt-4">
                <button disabled={updating} onClick={() => handleAction(selectedUser.id, "block")} className="bg-red-100 text-red-700 px-3 py-1 rounded-lg text-sm hover:bg-red-200">
                  🔒 Bloquear
                </button>
                <button disabled={updating} onClick={() => handleAction(selectedUser.id, "activate")} className="bg-green-100 text-green-700 px-3 py-1 rounded-lg text-sm hover:bg-green-200">
                  🔓 Ativar
                </button>
                <button disabled={updating} onClick={() => handleAction(selectedUser.id, "leave")} className="bg-yellow-100 text-yellow-700 px-3 py-1 rounded-lg text-sm hover:bg-yellow-200">
                  🏖️ Férias
                </button>
                <button disabled={updating} onClick={() => handleAction(selectedUser.id, "inactive")} className="bg-gray-100 text-gray-700 px-3 py-1 rounded-lg text-sm hover:bg-gray-200">
                  🚫 Demitir
                </button>
              </div>

              {/* Histórico */}
              <div className="mt-6">
                <h3 className="text-lg font-semibold mb-2">Histórico de Auditoria</h3>
                {loadingLogs ? (
                  <p className="text-gray-500 text-sm">Carregando histórico...</p>
                ) : logs.length === 0 ? (
                  <p className="text-gray-500 text-sm">Nenhuma ação registrada.</p>
                ) : (
                  <ul className="space-y-3">
                    {logs.map((log) => (
                      <li key={log.id} className="border border-gray-100 rounded-lg p-3 bg-gray-50 hover:bg-gray-100 transition">
                        <div className="flex items-center justify-between">
                          <span className={`text-xs px-2 py-1 rounded-full font-medium ${getActionColor(log.action)}`}>
                            {getActionLabel(log.action)}
                          </span>
                          <span className="text-xs text-gray-500">
                            {new Date(log.created_at).toLocaleString("pt-BR")}
                          </span>
                        </div>

                        {(log.old_value || log.new_value) && (
                          <div className="mt-2 text-sm text-gray-700">
                            <p><strong>Antes:</strong> {log.old_value || "—"}</p>
                            <p><strong>Depois:</strong> {log.new_value || "—"}</p>
                          </div>
                        )}

                        {log.reason && (
                          <p className="mt-2 text-sm italic text-gray-600">
                            💬 Motivo: {log.reason}
                          </p>
                        )}
                      </li>
                    ))}
                  </ul>
                )}
              </div>
            </div>
          </div>
        </div>
      )}
    </div>
  );
}
