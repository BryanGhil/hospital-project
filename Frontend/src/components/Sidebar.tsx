import { NavLink } from "react-router";
import { useAuthStore } from "../store/UseAuthStore";
import EmojiPeopleIcon from '@mui/icons-material/EmojiPeople';
import HomeIcon from '@mui/icons-material/Home';

export const Sidebar = () => {
  return (
    <aside className="w-32 md:w-48 h-screen bg-gray-800 text-white p-4 flex flex-col justify-between">
      <div className="flex flex-row gap-2 items-center">
        <img
          src="../../Logo.png"
          alt="app logo"
          className="w-5 h-5 rounded-sm"
        />
        <p className="text-2xl">Hospitality</p>
      </div>
      <nav className="flex flex-col gap-4 text-lg">
        <NavLink to="/" className="hover:bg-gray-700 p-2 rounded flex gap-2 items-center">
          <HomeIcon/> Home
        </NavLink>
        <NavLink to="/patient" className="hover:bg-gray-700 p-2 rounded flex gap-2 items-center" >
          <EmojiPeopleIcon/> Patient
        </NavLink>
      </nav>
      <NavLink to="/login" onClick={() => useAuthStore.getState().clearToken()} className="hover:bg-gray-700 p-2 rounded text-lg">Logout</NavLink>
    </aside>
  );
};
