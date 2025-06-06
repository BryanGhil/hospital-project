import { useEffect, useState } from "react";
import ax from "../../../api/axios";
import toast from "react-hot-toast";
import CircularProgress from "@mui/material/CircularProgress";

interface patientdata {
  patient_id: number;
  full_name: string;
  dob: string;
  gender: string;
  address: string;
  phone: string;
}

const defaultPatientData: patientdata = {
  patient_id: 0,
  full_name: "",
  dob: "",
  gender: "",
  address: "",
  phone: "",
};

export const ListPatient = () => {
  const [data, setData] = useState<patientdata[]>([defaultPatientData]);
  const [loading, setLoading] = useState(true);

  useEffect(() => {
    const abortController = new AbortController();

    const fetchData = async () => {
      try {
      const res = await ax.get("/patients", {
        signal: abortController.signal,
      });

      if (res.data?.data?.data) {
        setData(res.data.data.data);
      } else {
        throw new Error("Invalid response structure");
      }
    } catch (error) {
        if (!abortController.signal.aborted || error != null) { // Only handle non-canceled errors
          toast.error("Failed to load patient data");
        }
      } finally {
        if (!abortController.signal.aborted) {
          setLoading(false);
        }
      }
    };

    fetchData();

    return () => abortController.abort();
  }, []);

  if (loading) {
    return (
      <div className="flex justify-center items-center h-64">
        <CircularProgress style={{ color: 'black' }} /> {/* Black loader */}
      </div>
    );
  }

  
  return (
    <div className="overflow-x-auto rounded-lg border border-gray-200 shadow-sm">
      <table className="min-w-full divide-y divide-gray-200 mt-4">
        <thead className="bg-gray-50">
          <tr>
            <th className="px-6 py-3 text-left text-xs font-medium text-gray-500 uppercase tracking-wider">
              ID
            </th>
            <th className="px-6 py-3 text-left text-xs font-medium text-gray-500 uppercase tracking-wider">
              Name
            </th>
            <th className="px-6 py-3 text-left text-xs font-medium text-gray-500 uppercase tracking-wider">
              DOB
            </th>
            <th className="px-6 py-3 text-left text-xs font-medium text-gray-500 uppercase tracking-wider">
              Gender
            </th>
            <th className="px-6 py-3 text-left text-xs font-medium text-gray-500 uppercase tracking-wider">
              Address
            </th>
            <th className="px-6 py-3 text-left text-xs font-medium text-gray-500 uppercase tracking-wider">
              Phone
            </th>
          </tr>
        </thead>
        <tbody>
          {data.map((x, i) => (
            <tr
              key={x.patient_id}
              className={`
                ${i % 2 === 0 ? "bg-white" : "bg-gray-50"} 
                hover:bg-gray-100 cursor-pointer transition-colors duration-150`}
            >
              <td className="px-6 py-4 whitespace-nowrap text-sm ">
                {x.patient_id}
              </td>
              <td className="px-6 py-4 whitespace-nowrap text-sm ">
                {x.full_name}
              </td>
              <td className="px-6 py-4 whitespace-nowrap text-sm ">{x.dob}</td>
              <td className="px-6 py-4 whitespace-nowrap text-sm ">
                {x.gender}
              </td>
              <td className="px-6 py-4 whitespace-nowrap text-sm ">
                {x.address}
              </td>
              <td className="px-6 py-4 whitespace-nowrap text-sm ">
                {x.phone}
              </td>
            </tr>
          ))}
        </tbody>
      </table>
    </div>
  );
};
