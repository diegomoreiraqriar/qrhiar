export async function getUsers() {
  const res = await fetch("http://localhost:4000/third-parties", {
    method: "GET",
    headers: { "Content-Type": "application/json" },
  });
  if (!res.ok) throw new Error("Erro ao buscar usuários");
  return res.json();
}
