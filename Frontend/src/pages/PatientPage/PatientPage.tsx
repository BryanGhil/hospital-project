import { useState } from "react";
import { ListPatient } from "./components/ListPatient";
import { AddPatient } from "./components/AddPatient";

const patientMenu = ["List Patients", "Add Patient"] as const;

type SelectMenu = (typeof patientMenu)[number];

export const PatientPage = () => {
  const [activeMenu, setActiveMenu] = useState<SelectMenu>("List Patients");
  return (
    <div className="flex flex-col gap-4">
      <div>
        {patientMenu.map((menu) => (
          <button
            key={menu}
            onClick={() => setActiveMenu(menu)}
            className={`text-lg px-4 py-2 rounded-full mr-4 ${
              activeMenu === menu
                ? "bg-slate-900 text-white"
                : "bg-slate-300 hover:bg-slate-500 hover:text-white"
            }`}
          >
            {menu}
          </button>
        ))}
      </div>
      <div>
        {activeMenu === "List Patients" && <ListPatient/>}
        {activeMenu === "Add Patient" && <AddPatient/>}
      </div>
    </div>
  );
};
