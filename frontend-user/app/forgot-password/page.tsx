"use client"

import { useAuth } from "@/hooks/use-auth"
import ForgotPasswordForm from "@/components/forgot-password-form"

export default function ForgotPasswordPage() {
  useAuth(false) // Não requer autenticação, redireciona se já logado

  return <ForgotPasswordForm />
}
