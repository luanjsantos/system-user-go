import { create } from "zustand"
import { persist } from "zustand/middleware"
import { authService } from "./auth-service"

interface User {
  id: string
  email: string
  name: string
}

interface AuthState {
  user: User | null
  token: string | null
  isAuthenticated: boolean
  isLoading: boolean
  login: (email: string, password: string) => Promise<boolean>
  logout: () => void
  checkAuth: () => boolean
  setLoading: (loading: boolean) => void
}

// Função para verificar se o token está expirado
const isTokenExpired = (token: string): boolean => {
  try {
    const payload = JSON.parse(atob(token.split(".")[1]))
    const currentTime = Date.now() / 1000
    return payload.exp < currentTime
  } catch {
    return true
  }
}

export const useAuthStore = create<AuthState>()(
  persist(
    (set, get) => ({
      user: null,
      token: null,
      isAuthenticated: false,
      isLoading: false,

      login: async (email: string, password: string) => {
        set({ isLoading: true })

        try {
          const data = await authService.login(email, password)

          set({
            user: data.user,
            token: data.token,
            isAuthenticated: true,
            isLoading: false,
          })

          return true
        } catch (error) {
          set({ isLoading: false })
          console.error("Erro no login:", error)
          return false
        }
      },

      logout: () => {
        set({
          user: null,
          token: null,
          isAuthenticated: false,
        })
      },

      checkAuth: () => {
        const { token } = get()

        if (!token) {
          return false
        }

        if (isTokenExpired(token)) {
          get().logout()
          return false
        }

        return true
      },

      setLoading: (loading: boolean) => {
        set({ isLoading: loading })
      },
    }),
    {
      name: "auth-storage",
      partialize: (state) => ({
        user: state.user,
        token: state.token,
        isAuthenticated: state.isAuthenticated,
      }),
    },
  ),
)
