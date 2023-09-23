import {defineStore} from 'pinia'
import {ref} from 'vue'
import axios, {AxiosError} from 'axios'
import type {AuthUserType, UserRegistrationDto} from '@/types/user'
import type {Ref} from 'vue'
import {setAuthToken} from '@/lib/http'
import type {ErrorResponse} from '@/types/validationErrors'
import {makeErrorResponse} from '@/types/validationErrors'

type AuthStoreReturn = {
    user: Ref<AuthUserType | null>
    authErrors: Ref<ErrorResponse | null>
    register: (userReg: UserRegistrationDto) => void
}

export const useAuthStore = defineStore('auth', (): AuthStoreReturn => {
    const user = ref<AuthUserType | null>(null)
    const authErrors = ref<ErrorResponse>({message: null, errors: null})

    const register = async (userReg: UserRegistrationDto) => {
        try {
            const {data} = await axios.post<AuthUserType>('/api/register', userReg)
            setAuthToken(data.token)
        } catch (error) {
            if (!(error instanceof AxiosError)) {
                console.log(error)
                return
            }
            authErrors.value = makeErrorResponse(error?.response?.data)
        }
    }

    return {
        user,
        authErrors,
        register
    }
})
