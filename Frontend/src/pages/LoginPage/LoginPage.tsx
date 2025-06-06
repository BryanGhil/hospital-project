import {
  Button,
  CircularProgress,
  FormControl,
  FormHelperText,
  IconButton,
  InputAdornment,
  InputLabel,
  OutlinedInput,
  Typography,
} from "@mui/material";
import { Visibility, VisibilityOff } from "@mui/icons-material";
import { useCallback, useState } from "react";
import { z } from "zod";
import ax from "../../api/axios";
import { useAuthStore } from "../../store/UseAuthStore";
import { useNavigate } from "react-router";
import toast from "react-hot-toast";
import { handleAxiosError } from "../../utils/axiosErrorHandler";

const schema = z.object({
  email: z
    .string()
    .min(1, "Email is required")
    .email("Please enter a valid email"),
  password: z.string().min(1, "Password is required"),
});

type FormData = {
  email: string;
  password: string;
};

type FormErrors = Partial<Record<keyof FormData, string>>;

export const LoginPage = () => {
  const [showPassword, setShowPassword] = useState(false);
  const [form, setForm] = useState({ email: "", password: "" });
  const [errors, setErrors] = useState<FormErrors>({});
  const [loading, setLoading] = useState(false);

  const navigate = useNavigate();
  const setToken = useAuthStore((state) => state.setToken);

  const handleClickShowPassword = useCallback(
    () => setShowPassword((prev) => !prev),
    []
  );

  const handleChange = (e: React.ChangeEvent<HTMLInputElement>) => {
    const { name, value } = e.target;
    setForm((prev) => ({ ...prev, [name]: value }));

    // Real-time validation after first submission attempt
    if (Object.keys(errors).length > 0) {
      const result = schema.safeParse({ ...form, [name]: value });
      const newErrors: FormErrors = {};

      if (!result.success) {
        result.error.errors.forEach((err) => {
          const field = err.path[0] as keyof FormData;
          newErrors[field] = err.message;
        });
      }

      setErrors((prev) => {
        // Clear error for the current field if it's valid
        const updatedErrors = { ...prev };
        if (!newErrors[name as keyof FormData]) {
          delete updatedErrors[name as keyof FormData];
        }
        return { ...updatedErrors, ...newErrors };
      });
    }
  };

  const handleSubmit = async (e: React.FormEvent) => {
    e.preventDefault();
    setLoading(true);

    // Validate form
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
      const response = await ax.post("/login", {
        email: form.email,
        password: form.password,
      });

      setToken(response.data.data.token);
      toast.success("Login successful!");
      navigate("/");
    } catch (err) {
      const errorMessage = handleAxiosError(err);

      toast.error(errorMessage);
    } finally {
      setLoading(false);
    }
  };

  return (
    <div className="flex flex-row md:flex-row w-screen h-screen bg-gradient-to-b from-gray-200 to-gray-300">
      <div className="w-full md:w-1/2 flex flex-col justify-center px-4 md:px-20 gap-6">
        <Typography variant="h3">Welcome back!</Typography>
        <Typography variant="body1">Please enter your details</Typography>
        <form onSubmit={handleSubmit} className="flex flex-col gap-6">
          <FormControl variant="outlined" fullWidth error={!!errors.email}>
            <InputLabel htmlFor="email">Email</InputLabel>
            <OutlinedInput
              id="email"
              name="email"
              value={form.email}
              onChange={handleChange}
              type="email"
              label="Email"
              disabled={loading}
            />
            {errors.email && <FormHelperText>{errors.email}</FormHelperText>}
          </FormControl>
          <FormControl variant="outlined" fullWidth error={!!errors.password}>
            <InputLabel htmlFor="password">Password</InputLabel>
            <OutlinedInput
              id="password"
              name="password"
              value={form.password}
              onChange={handleChange}
              type={showPassword ? "text" : "password"}
              disabled={loading}
              endAdornment={
                <InputAdornment position="end">
                  <IconButton
                    aria-label={
                      showPassword ? "Hide password" : "Show password"
                    }
                    onClick={handleClickShowPassword}
                    edge="end"
                    disabled={loading}
                  >
                    {showPassword ? <VisibilityOff /> : <Visibility />}
                  </IconButton>
                </InputAdornment>
              }
              label="Password"
            />
            {errors.password && (
              <FormHelperText>{errors.password}</FormHelperText>
            )}
          </FormControl>
          <Button
            startIcon={loading ? <CircularProgress size={20} sx={{color: "black"}} /> : undefined}
            sx={{
              padding: "12px 24px",
              backgroundColor: "black",
              color: "white",
              "&:hover": {
                backgroundColor: "gray",
              },
            }}
            variant="contained"
            fullWidth
            disabled={loading}
            type="submit"
          >
            {loading ? "Signing in..." : "Submit"}
          </Button>
        </form>
      </div>
      <img
        src="/cover.jpg"
        alt="Doctor with a stethoscope"
        className="hidden md:block w-1/2 m-4 rounded-3xl object-cover"
      />
    </div>
  );
};
