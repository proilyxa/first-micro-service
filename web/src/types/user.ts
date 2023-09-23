export type UserRegistrationDto = {
    firstName: string | null
    lastName: string | null
    email: string | null
    password: string | null
    passwordConfirm: string | null
}

export type AuthUserType = {
    id: number
    firstName: string
    lastName?: string
    email: string
    token: string
}
