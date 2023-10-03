import { useState } from "react";

class ErrorResponse {
  message: string | null = null;
  errors: Map<string, string[]> | null = null;

  constructor(dataResponse: any) {
    this.makeErrors(dataResponse);
  }

  private makeErrors(dataResponse: any): void {
    this.message = dataResponse?.message ?? null;

    if (dataResponse.errors instanceof Object) {
      this.errors = new Map<string, string[]>();
      for (const [key, value] of Object.entries(dataResponse.errors)) {
        this.errors.set(key, value as string[]);
      }
    }
  }

  public getErrorsByKey(key: string): string[] | null {
    return this.errors?.get(key) ?? null;
  }
}

const useErrorResponse = (): [ErrorResponse | null, (error: any) => void] => {
  const [errorResponse, setErrorResponse] = useState<ErrorResponse | null>(
    null,
  );

  const setErrorFromHttpError = (error: any): void => {
    setErrorResponse(new ErrorResponse(error?.response?.data));
  };

  return [errorResponse, setErrorFromHttpError];
};

export { ErrorResponse, useErrorResponse };
