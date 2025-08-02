import { useAuthStore } from "./auth-store"

class AuthService {
  private baseURL = "/api"

  // Interceptor para adicionar token nas requisições
  private async request(url: string, options: RequestInit = {}) {
    const { token } = useAuthStore.getState()

    const headers: Record<string, string> = {
      "Content-Type": "application/json",
    }

    // Adicionar headers customizados se existirem
    if (options.headers) {
      Object.entries(options.headers).forEach(([key, value]) => {
        if (typeof value === 'string') {
          headers[key] = value
        }
      })
    }

    if (token && !this.isTokenExpired(token)) {
      headers["Authorization"] = `Bearer ${token}`
    }

    const response = await fetch(`${this.baseURL}${url}`, {
      ...options,
      headers,
    })

    // Se receber 401, token expirou
    if (response.status === 401) {
      useAuthStore.getState().logout()
      window.location.href = "/login"
    }

    return response
  }

  private isTokenExpired(token: string): boolean {
    try {
      const payload = JSON.parse(atob(token.split(".")[1]))
      const currentTime = Date.now() / 1000
      return payload.exp < currentTime
    } catch {
      return true
    }
  }

  async login(email: string, password: string) {
    const response = await this.request("/auth/login", {
      method: "POST",
      body: JSON.stringify({ email, password }),
    })

    if (!response.ok) {
      const errorData = await response.json().catch(() => ({}))
      throw new Error(errorData.message || "Credenciais inválidas")
    }

    return response.json()
  }

  async getProfile() {
    const response = await this.request("/profile")

    if (!response.ok) {
      throw new Error("Erro ao buscar perfil")
    }

    return response.json()
  }

  async refreshToken() {
    const response = await this.request("/auth/refresh", {
      method: "POST",
    })

    if (!response.ok) {
      throw new Error("Erro ao renovar token")
    }

    return response.json()
  }

  async forgotPassword(email: string) {
    const response = await fetch(`${this.baseURL}/auth/reset-password`, {
      method: "POST",
      headers: {
        "Content-Type": "application/json",
      },
      body: JSON.stringify({ email }),
    })

    if (!response.ok) {
      const errorData = await response.json().catch(() => ({}))
      throw new Error(errorData.error || "Erro ao enviar email de recuperação")
    }

    return response.json()
  }
}

export const authService = new AuthService()
