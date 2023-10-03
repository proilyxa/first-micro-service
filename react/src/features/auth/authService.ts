import { AuthUserType, UserRegistrationDto } from "@/@types/user.ts";
import { instance as axios } from "@/lib/http.ts";
import { FormErrorException } from "@/services/form-error-response/formErrorException.ts";
import { ErrorResponse } from "@/services/form-error-response/errorResponse.ts";
import { AxiosError } from "axios";

export const AuthService = {
  register: async (
    userDto: UserRegistrationDto,
  ): Promise<AuthUserType | never> => {
    try {
      const { data } = await axios.post<AuthUserType>(
        "/auth/register",
        userDto,
      );
      return data;
    } catch (error) {
      if (!(error instanceof AxiosError)) {
        throw error;
      }

      throw new FormErrorException(new ErrorResponse(error.response?.data));
    }
  },
};
