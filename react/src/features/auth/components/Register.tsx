import React, { FC, useState } from "react";
import styles from "../../../components/ui/input/auth.module.css";
import { Link } from "react-router-dom";
import { UserRegistrationDto } from "@/@types/user.ts";
import { useErrorResponse } from "@/services/form-error-response/errorResponse.ts";
import InputField from "@/components/ui/input/InputField.tsx";
import { useMutation } from "@tanstack/react-query";
import { instance as axios } from "@/lib/http.ts";
import { AxiosError } from "axios";

const Register: FC = () => {
  console.log("render");

  const [form, setForm] = useState<UserRegistrationDto>({
    firstName: "",
    lastName: "",
    email: "",
    password: "",
    passwordConfirm: "",
  });

  const [errorResponse, setErrorResponse] = useErrorResponse();

  const mutation = useMutation({
    mutationFn: (userDto: UserRegistrationDto) => {
      return axios.post("/auth/register", userDto);
    },
    onError: (error: AxiosError) => {
      setErrorResponse(error);
    },
  });
  console.log(mutation);

  const onInputChange = (e: React.ChangeEvent<HTMLInputElement>): void => {
    setForm({
      ...form,
      [e.target.name]: e.target.value,
    });
  };

  const onSubmit = (e: React.FormEvent<HTMLFormElement>): void => {
    e.preventDefault();
    mutation.mutate(form);
  };

  return (
    <div className={"flex flex-col gap-4"}>
      <div className={"flex flex-col items-center gap-2"}>
        <h2 className={"text-4xl "}>Register</h2>
        <Link to={"/login"} className={"text-xs text-blue-700"}>
          already have account?
        </Link>
      </div>

      <form className={""} onSubmit={onSubmit}>
        <InputField
          label={"First Name"}
          id="firstName"
          name="firstName"
          type="text"
          onChange={onInputChange}
          errors={errorResponse?.getErrorsByKey("firstName")}
        ></InputField>

        <InputField
          label={"Last Name"}
          id="lastName"
          name="lastName"
          type="text"
          onChange={onInputChange}
          errors={errorResponse?.getErrorsByKey("lastName")}
        ></InputField>

        <InputField
          label={"Email"}
          id="email"
          name="email"
          type="text"
          onChange={onInputChange}
          errors={errorResponse?.getErrorsByKey("email")}
        ></InputField>

        <InputField
          label={"Password"}
          id="password"
          name="password"
          type="password"
          onChange={onInputChange}
          errors={errorResponse?.getErrorsByKey("password")}
        ></InputField>

        <InputField
          label={"Password Confirmation"}
          id="passwordConfirm"
          name="passwordConfirm"
          type="password"
          onChange={onInputChange}
          errors={errorResponse?.getErrorsByKey("passwordConfirm")}
        ></InputField>

        <button type="submit" className={`mt-2 ${styles.btn}`}>
          Register
        </button>
      </form>
    </div>
  );
};

export default Register;
