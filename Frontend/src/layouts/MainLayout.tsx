import { Outlet } from "react-router"
import { Sidebar } from "../components/Sidebar"

export const MainLayout = () => {
  return (
    <div className="flex h-screen">
        <Sidebar/>
        <main className="flex-1 bg-gradient-to-b from-gray-200 to-gray-300 p-6 overflow-auto">
            <Outlet/>
        </main>

    </div>
  )
}
