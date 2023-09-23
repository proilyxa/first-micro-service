export type ErrorResponse = {
    message: string | null
    errors: Map<string, string[]> | null
}


export const makeErrorResponse = (dataResponse: any): ErrorResponse => {
    const resp: ErrorResponse = {
        message: dataResponse.message ?? null,
        errors: null,
    }

    if (dataResponse.errors instanceof Object) {
        resp.errors = new Map<string, string[]>()
        for (const [key, value] of Object.entries(dataResponse.errors)) {
            resp.errors.set(key, value as string[])
        }
    }

    return resp;
}