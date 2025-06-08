export interface patientdata {
  patient_id: number;
  full_name: string;
  dob: string;
  gender: string;
  address: string;
  phone: string;
}

export const defaultPatientData: patientdata = {
  patient_id: 0,
  full_name: "",
  dob: "",
  gender: "",
  address: "",
  phone: "",
};