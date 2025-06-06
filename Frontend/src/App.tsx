import { Navigate, Route, Routes } from "react-router";
import { MainLayout } from "./layouts/MainLayout";
import { Homepage } from "./pages/Homepage/Homepage";
import { PatientPage } from "./pages/PatientPage/PatientPage";
import { LoginPage } from "./pages/LoginPage/LoginPage";
import { ProtectedRoute } from "./layouts/ProtectedRoute";
import { PublicRoute } from "./layouts/PublicRoute";

export function App() {
  return (
    <Routes>
      <Route element={<PublicRoute />}>
        <Route path="/login" element={<LoginPage />} />
      </Route>

      <Route element={<ProtectedRoute />}>
        <Route path="/" element={<MainLayout />}>
          <Route index element={<Homepage />}></Route>
          <Route path="patient" element={<PatientPage />}></Route>
        </Route>
      </Route>

      <Route path="*" element={<Navigate to="/" replace />} />
    </Routes>
  );
}
