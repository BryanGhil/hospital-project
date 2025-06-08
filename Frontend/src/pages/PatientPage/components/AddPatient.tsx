import {
  Button,
  CircularProgress,
  FormControl,
  FormHelperText,
  InputLabel,
  OutlinedInput,
  Select,
  MenuItem,
  type SelectChangeEvent,
} from "@mui/material";
import { useState } from "react";
import z from "zod";
import { DatePicker } from "@mui/x-date-pickers/DatePicker";
import { LocalizationProvider } from "@mui/x-date-pickers/LocalizationProvider";
import { AdapterDateFns } from "@mui/x-date-pickers/AdapterDateFns";
import { format } from "date-fns";
import ax from "../../../api/axios";
import toast from "react-hot-toast";
import { handleAxiosError } from "../../../utils/axiosErrorHandler";

const schema = z.object({
  name: z.string().min(1, "Name is required"),
  dob: z.string().min(1, "Date of Birth is required"),
  gender: z.string().min(1, "Gender is required"),
  address: z.string().min(1, "Address is required"),
  phone: z.string().min(1, "Phone is required"),
});

type FormData = {
  name: string;
  dob: string;
  gender: string;
  address: string;
  phone: string;
};

const defaultPatientData: FormData = {
  name: "",
  dob: "",
  gender: "",
  address: "",
  phone: "",
};

type FormErrors = Partial<Record<keyof FormData, string>>;

export const AddPatient = () => {
  const [form, setForm] = useState<FormData>(defaultPatientData);
  const [errors, setErrors] = useState<FormErrors>({});
  const [loading, setLoading] = useState(false);

  const validateField = (fieldName: keyof FormData, value: string) => {
    const fieldSchema = schema.shape[fieldName];
    const result = fieldSchema.safeParse(value);

    if (!result.success) {
      const error = result.error.errors[0];
      setErrors((prev) => ({ ...prev, [fieldName]: error.message }));
    } else {
      setErrors((prev) => {
        const newErrors = { ...prev };
        delete newErrors[fieldName];
        return newErrors;
      });
    }
  };

  const handleInputChange = (e: React.ChangeEvent<HTMLInputElement>) => {
    const { name, value } = e.target;
    setForm((prev) => ({ ...prev, [name]: value }));
    validateField(name as keyof FormData, value);
  };

  const handleSelectChange = (e: SelectChangeEvent<string>) => {
    const name = e.target.name as keyof FormData;
    const value = e.target.value;
    setForm((prev) => ({ ...prev, [name]: value }));
    validateField(name, value);
  };

  const handleSubmit = async (e: React.FormEvent) => {
    e.preventDefault();
    setLoading(true);

    const result = schema.safeParse(form);
    if (!result.success) {
      const fieldErrors: FormErrors = {};
      result.error.errors.forEach((err) => {
        const field = err.path[0] as keyof FormData;
        fieldErrors[field] = err.message;
      });
      setErrors(fieldErrors);
      setLoading(false);
      return;
    }

    try {
      const response = await ax.post("/patients", {
        full_name: form.name,
        dob: form.dob,
        gender: form.gender,
        address: form.address,
        phone: form.phone,
      });

      if (response.status === 201) {
        toast.success("Patient Data Succesfully Added");
        setForm(defaultPatientData);
      } else {
        throw new Error("error occured");
      }
    } catch (err) {
      const errorMessage = handleAxiosError(err);

      toast.error(errorMessage);
    } finally {
      setLoading(false);
    }
  };

  return (
    <div className="flex justify-center mt-4">
      <form
        onSubmit={handleSubmit}
        className="p-4 bg-white rounded-lg w-full flex flex-col justify-center justify-items-center gap-4"
      >
        <p className="text-3xl font-bold">Add Patient</p>
        <div className="flex flex-row gap-4">
          <div className="flex flex-col gap-4 w-1/2">
            <FormControl variant="outlined" fullWidth error={!!errors.name}>
              <InputLabel htmlFor="name">Name</InputLabel>
              <OutlinedInput
                id="name"
                name="name"
                value={form.name}
                onChange={handleInputChange}
                type="text"
                label="name"
                disabled={loading}
              />
              {errors.name && <FormHelperText>{errors.name}</FormHelperText>}
            </FormControl>

            <LocalizationProvider dateAdapter={AdapterDateFns}>
              <DatePicker
                label="Date of Birth"
                value={form.dob ? new Date(form.dob) : null}
                onChange={(newValue: Date | null) => {
                  const dobString = newValue
                    ? format(newValue, "yyyy-MM-dd")
                    : "";
                  setForm((prev) => ({ ...prev, dob: dobString }));
                  validateField("dob", dobString);
                }}
                format="dd/MM/yyyy"
                disabled={loading}
                slotProps={{
                  textField: {
                    fullWidth: true,
                    error: !!errors.dob,
                    helperText: errors.dob,
                  },
                }}
              />
            </LocalizationProvider>

            <FormControl variant="outlined" fullWidth error={!!errors.gender}>
              <InputLabel htmlFor="gender">Gender</InputLabel>
              <Select
                id="gender"
                name="gender"
                value={form.gender}
                onChange={handleSelectChange}
                label="gender"
                disabled={loading}
              >
                <MenuItem value="">
                  <em>Select Gender</em>
                </MenuItem>
                <MenuItem value="Male">Male</MenuItem>
                <MenuItem value="Female">Female</MenuItem>
              </Select>
              {errors.gender && (
                <FormHelperText>{errors.gender}</FormHelperText>
              )}
            </FormControl>
          </div>

          <div className="flex flex-col gap-4 w-1/2">
            <FormControl variant="outlined" fullWidth error={!!errors.address}>
              <InputLabel htmlFor="address">Address</InputLabel>
              <OutlinedInput
                id="address"
                name="address"
                value={form.address}
                onChange={handleInputChange}
                type="text"
                label="address"
                disabled={loading}
              />
              {errors.address && (
                <FormHelperText>{errors.address}</FormHelperText>
              )}
            </FormControl>

            <FormControl variant="outlined" fullWidth error={!!errors.phone}>
              <InputLabel htmlFor="phone">Phone</InputLabel>
              <OutlinedInput
                id="phone"
                name="phone"
                value={form.phone}
                onChange={handleInputChange}
                type="tel"
                label="phone"
                disabled={loading}
              />
              {errors.phone && <FormHelperText>{errors.phone}</FormHelperText>}
            </FormControl>
          </div>
        </div>

        <Button
          startIcon={
            loading ? (
              <CircularProgress size={20} sx={{ color: "black" }} />
            ) : undefined
          }
          sx={{
            padding: "12px 24px",
            backgroundColor: "black",
            color: "white",
            "&:hover": {
              backgroundColor: "gray",
            },
          }}
          variant="contained"
          disabled={loading}
          type="submit"
          className="w-1/4"
        >
          {loading ? "" : "Submit"}
        </Button>
      </form>
    </div>
  );
};
