import axios from "axios";

const API_BASE_URL = import.meta.env.VITE_API_URL || "http://localhost:4000";

export async function login(email: string, password: string) {
  const response = await axios.post(`${API_BASE_URL}/auth/login`, { email, password });
  const { token } = response.data;
  localStorage.setItem("token", token);
  return token;
}

export function getToken() {
  return localStorage.getItem("token");
}

export function logout() {
  localStorage.removeItem("token");
}
