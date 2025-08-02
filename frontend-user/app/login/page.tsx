"use client"

import { useAuth } from "@/hooks/use-auth"
import LoginForm from "@/components/login-form"

export default function LoginPage() {
  useAuth(false) // Não requer autenticação, redireciona se já logado

  return <LoginForm />
}
