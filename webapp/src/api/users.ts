import api from "./client";

export async function getUsers() {
  const res = await api.get("/api/third-parties");
  return res.data;
}
