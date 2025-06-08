import { useEffect, useState } from "react";
import ax from "../../../api/axios";
import CircularProgress from "@mui/material/CircularProgress";
import { useNavigate } from "react-router";
import { defaultPatientData, type patientdata } from "../types/Patient";

export const ListPatient = () => {
  const [data, setData] = useState<patientdata[]>([defaultPatientData]);
  const [loading, setLoading] = useState(true);
  const [error, setError] = useState<string | null>(null);

  const navigate = useNavigate();

  useEffect(() => {
    const abortController = new AbortController();

    const fetchData = async () => {
      try {
      const res = await ax.get("/patients", {
        signal: abortController.signal,
      });

      if (res.data?.data?.data) {
        setData(res.data.data.data);
        setError(null)
      } else {
        throw new Error("Invalid response structure");
      }
    } catch (err) {
        if (!abortController.signal.aborted || err != null) { 
          setError("Failed to load patient data");
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
        <CircularProgress style={{ color: 'black' }} />
      </div>
    );
  }

  if (error) {
    return (
      <div className="text-center py-8 text-black text-xl font-bold">
        Please try again later
      </div>
    );
  }

  
  return (
    <div className="overflow-x-auto rounded-lg border border-gray-200 shadow-sm">
      <table className="min-w-full divide-y divide-gray-200 mt-4">
        <thead className="bg-gray-50">
          <tr>
            <th className="px-6 py-3 text-left text-xs font-bold text-gray-500 uppercase tracking-wider">
              ID
            </th>
            <th className="px-6 py-3 text-left text-xs font-bold text-gray-500 uppercase tracking-wider">
              Name
            </th>
            <th className="px-6 py-3 text-left text-xs font-bold text-gray-500 uppercase tracking-wider">
              DOB
            </th>
            <th className="px-6 py-3 text-left text-xs font-bold text-gray-500 uppercase tracking-wider">
              Gender
            </th>
            <th className="px-6 py-3 text-left text-xs font-bold text-gray-500 uppercase tracking-wider">
              Address
            </th>
            <th className="px-6 py-3 text-left text-xs font-bold text-gray-500 uppercase tracking-wider">
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
              onClick={() => navigate(`/patient/${x.patient_id}`)}
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
