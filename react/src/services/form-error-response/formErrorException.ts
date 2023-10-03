import { ErrorResponse } from "@/services/form-error-response/errorResponse.ts";

export class FormErrorException extends Error {
  public errorResponse: ErrorResponse

  constructor(errorResponse: ErrorResponse) {
    super();
    this.errorResponse = errorResponse;
  }

  public getErrorResponse(): ErrorResponse {
    return this.errorResponse;
  }
}

