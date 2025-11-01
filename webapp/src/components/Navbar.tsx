import { branding } from "../config/branding";

export default function Navbar() {
  return (
    <nav className="bg-white shadow flex items-center justify-between px-6 py-3">
      <div className="flex items-center gap-3">
        <img src={branding.logoUrl} alt="Logo" className="h-8 w-auto" />
        <h1 className="text-xl font-semibold text-gray-800">
          {branding.companyName}
        </h1>
      </div>
      <div className="text-sm text-gray-500">
        <span>Identidade: RH de Terceiros</span>
      </div>
    </nav>
  );
}
