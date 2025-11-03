import { useEffect, useState } from "react";
import toast from "react-hot-toast";
import api from "../api/client"; // ✅ usa axios configurado com token e baseURL

interface Props {
  onClose: () => void;
  onUserCreated: () => void;
}

interface Company {
  id: string;
  name: string;
}

interface Manager {
  id: string;
  name: string;
}

export default function CreateUserModal({ onClose, onUserCreated }: Props) {
  const [form, setForm] = useState({
    name: "",
    email: "",
    cpf: "",
    position: "",
    company_id: "",
    manager_id: "",
  });
  const [companies, setCompanies] = useState<Company[]>([]);
  const [managers, setManagers] = useState<Manager[]>([]);
  const [loading, setLoading] = useState(false);
  const [error, setError] = useState("");

  // 🔹 Carregar empresas e gestores (agora usando o axios "api")
  useEffect(() => {
    const fetchData = async () => {
      try {
        const [companiesRes, managersRes] = await Promise.all([
          api.get("/api/companies"),
          api.get("/api/third-parties"),
        ]);

        setCompanies(companiesRes.data || []);
        setManagers(managersRes.data || []);
      } catch (err) {
        console.error("Erro ao carregar dados:", err);
        toast.error("Erro ao carregar empresas ou gestores");
      }
    };
    fetchData();
  }, []);

  // 🔹 Atualizar campos do formulário
  const handleChange = (e: React.ChangeEvent<HTMLInputElement | HTMLSelectElement>) => {
    setForm({ ...form, [e.target.name]: e.target.value });
  };

  // 🔹 Submeter formulário (criar usuário)
  const handleSubmit = async (e: React.FormEvent) => {
    e.preventDefault();
    setLoading(true);
    setError("");

    try {
      const res = await api.post("/api/third-parties", form);

      if (!res || res.status !== 201) {
        throw new Error("Erro ao criar usuário");
      }

      toast.success("Usuário criado com sucesso!");
      onUserCreated();
      onClose();
    } catch (err: any) {
      console.error("Erro ao criar usuário:", err);
      setError(err.message || "Erro inesperado");
      toast.error("Falha ao criar usuário");
    } finally {
      setLoading(false);
    }
  };

  return (
    <div className="fixed inset-0 bg-black bg-opacity-50 flex justify-center items-center p-4 z-50">
      <div className="bg-white rounded-xl shadow-lg w-full max-w-lg p-6">
        <h2 className="text-xl font-semibold mb-4 text-gray-900">Novo Usuário</h2>

        {error && <p className="text-red-500 text-sm mb-2">{error}</p>}

        <form onSubmit={handleSubmit} className="space-y-4">
          <input
            name="name"
            placeholder="Nome completo"
            value={form.name}
            onChange={handleChange}
            className="w-full p-2 border rounded-md"
            required
          />

          <input
            name="email"
            placeholder="E-mail"
            type="email"
            value={form.email}
            onChange={handleChange}
            className="w-full p-2 border rounded-md"
          />

          <input
            name="cpf"
            placeholder="CPF"
            value={form.cpf}
            onChange={handleChange}
            className="w-full p-2 border rounded-md"
          />

          <input
            name="position"
            placeholder="Cargo"
            value={form.position}
            onChange={handleChange}
            className="w-full p-2 border rounded-md"
          />

          {/* 🔹 Seleção de Empresa */}
          <div>
            <label className="block text-sm font-medium text-gray-700 mb-1">
              Empresa
            </label>
            <select
              name="company_id"
              value={form.company_id}
              onChange={handleChange}
              required
              className="w-full p-2 border rounded-md bg-white"
            >
              <option value="">Selecione uma empresa</option>
              {companies.map((c) => (
                <option key={c.id} value={c.id}>
                  {c.name}
                </option>
              ))}
            </select>
          </div>

          {/* 🔹 Seleção de Gestor */}
          <div>
            <label className="block text-sm font-medium text-gray-700 mb-1">
              Gestor
            </label>
            <select
              name="manager_id"
              value={form.manager_id}
              onChange={handleChange}
              required
              className="w-full p-2 border rounded-md bg-white"
            >
              <option value="">Selecione um gestor</option>
              {managers.map((m) => (
                <option key={m.id} value={m.id}>
                  {m.name}
                </option>
              ))}
            </select>
          </div>

          {/* 🔘 Botões */}
          <div className="flex justify-end space-x-3 mt-4">
            <button
              type="button"
              onClick={onClose}
              className="px-4 py-2 bg-gray-200 hover:bg-gray-300 rounded-md"
            >
              Cancelar
            </button>
            <button
              type="submit"
              disabled={loading}
              className="px-4 py-2 bg-indigo-600 text-white rounded-md hover:bg-indigo-700 disabled:opacity-50"
            >
              {loading ? "Salvando..." : "Criar Usuário"}
            </button>
          </div>
        </form>
      </div>
    </div>
  );
}
