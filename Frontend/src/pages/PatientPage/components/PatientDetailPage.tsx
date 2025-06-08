import { useParams } from "react-router";
import { defaultPatientData, type patientdata } from "../types/Patient";
import { useEffect, useState } from "react";
import CircularProgress from "@mui/material/CircularProgress";
import ax from "../../../api/axios";

export const PatientDetailPage = () => {
  const { id } = useParams();
  const [data, setData] = useState<patientdata>(defaultPatientData);
  const [loading, setLoading] = useState(true);
  const [error, setError] = useState<string | null>(null);

  useEffect(() => {
    const abortController = new AbortController();

    const fetchData = async () => {
      try {
        const res = await ax.get(`/patients/${id}`, {
          signal: abortController.signal,
        });

        if (res.data?.data) {
          setData(res.data.data);
          setError(null);
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
  }, [id]);

  if (loading) {
    return (
      <div className="flex justify-center items-center h-64">
        <CircularProgress style={{ color: "black" }} />
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
    <div className="flex flex-col w-full bg-white rounded-lg p-4 gap-6">
      <p className="text-3xl font-bold">Patient Detail</p>
      <div className="flex flex-row gap-4">
        <div className="w-1/2 flex flex-col gap-4">
          <div>
            <p className="text-xl font-bold">Name</p>
            <p className="text-lg">{data.full_name}</p>
          </div>
          <div>
            <p className="text-xl font-bold">Date of Birth</p>
            <p className="text-lg">{data.dob}</p>
          </div>
          <div>
            <p className="text-xl font-bold">Gender</p>
            <p className="text-lg">{data.gender}</p>
          </div>
        </div>
        <div className="w-1/2 flex flex-col gap-4">
          <div>
            <p className="text-xl font-bold">Address</p>
            <p className="text-lg">{data.address}</p>
          </div>
          <div>
            <p className="text-xl font-bold">Phone</p>
            <p className="text-lg">{data.phone}</p>
          </div>
        </div>
      </div>
    </div>
  );
};
