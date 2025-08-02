"use client"

import type React from "react"

import { useState } from "react"
import { useRouter } from "next/navigation"
import { authService } from "@/lib/auth-service"
import { Button } from "@/components/ui/button"
import { Input } from "@/components/ui/input"
import { Label } from "@/components/ui/label"
import { Alert, AlertDescription } from "@/components/ui/alert"
import { Loader2, Mail, ArrowLeft, CheckCircle } from "lucide-react"

export default function ForgotPasswordForm() {
  const [email, setEmail] = useState("")
  const [error, setError] = useState("")
  const [isLoading, setIsLoading] = useState(false)
  const [isSuccess, setIsSuccess] = useState(false)

  const router = useRouter()

  const handleSubmit = async (e: React.FormEvent) => {
    e.preventDefault()
    setError("")
    setIsLoading(true)

    if (!email) {
      setError("Por favor, digite seu email")
      setIsLoading(false)
      return
    }

    try {
      await authService.forgotPassword(email)
      setIsSuccess(true)
    } catch (err: unknown) {
      const errorMessage = err instanceof Error ? err.message : "Erro ao enviar email. Tente novamente."
      setError(errorMessage)
    } finally {
      setIsLoading(false)
    }
  }

  if (isSuccess) {
    return (
      <div className="min-h-screen flex">
        {/* Left Side - Image (70%) */}
        <div className="hidden lg:flex lg:w-[70%] bg-gradient-to-br from-green-600 via-emerald-600 to-teal-700 relative overflow-hidden">
          <div className="absolute inset-0 bg-black/20" />
          <div className="relative z-10 flex flex-col justify-center items-center p-12 text-white w-full">
            <div className="max-w-2xl text-center">
              <h1 className="text-5xl font-bold mb-8">Email Enviado!</h1>
              <p className="text-2xl opacity-90 mb-12 leading-relaxed">
                Verifique sua caixa de entrada e siga as instruções para redefinir sua senha
              </p>

              {/* Success illustration placeholder */}
              <div className="w-[500px] h-[400px] bg-white/10 rounded-3xl flex items-center justify-center backdrop-blur-sm mx-auto">
                <div className="text-center">
                  <div className="w-40 h-40 bg-white/20 rounded-full mx-auto mb-6 flex items-center justify-center">
                    <CheckCircle className="w-20 h-20 text-white/80" />
                  </div>
                  <p className="text-white/70 text-lg">Email de recuperação enviado</p>
                  <p className="text-white/50 text-sm mt-2">Verifique sua caixa de entrada</p>
                </div>
              </div>
            </div>
          </div>

          {/* Decorative elements */}
          <div className="absolute top-0 right-0 w-60 h-60 bg-white/10 rounded-full -translate-y-32 translate-x-32" />
          <div className="absolute bottom-0 left-0 w-48 h-48 bg-white/10 rounded-full translate-y-24 -translate-x-24" />
        </div>

        {/* Right Side - Success Message */}
        <div className="flex-1 lg:w-[30%] flex items-center justify-center p-6 bg-gray-50">
          <div className="w-full max-w-sm">
            <div className="text-center mb-8">
              <div className="inline-flex items-center justify-center w-16 h-16 bg-gradient-to-r from-green-600 to-emerald-600 rounded-2xl mb-6">
                <CheckCircle className="w-8 h-8 text-white" />
              </div>
              <h2 className="text-2xl font-bold text-gray-900 mb-2">Email Enviado!</h2>
              <p className="text-gray-600 text-sm">Verifique sua caixa de entrada</p>
            </div>

            <div className="bg-white rounded-2xl shadow-xl p-6 border border-gray-100">
              <div className="text-center space-y-4">
                <div className="p-4 bg-green-50 rounded-xl border border-green-200">
                  <p className="text-sm text-green-800 font-medium mb-2">Email enviado para:</p>
                  <p className="text-green-700 font-semibold">{email}</p>
                </div>

                <div className="text-left space-y-2 text-sm text-gray-600">
                  <p className="font-medium text-gray-700">Próximos passos:</p>
                  <ul className="space-y-1 list-disc list-inside">
                    <li>Verifique sua caixa de entrada</li>
                    <li>Clique no link do email</li>
                    <li>Defina uma nova senha</li>
                  </ul>
                </div>

                <div className="pt-4 space-y-3">
                  <Button
                    onClick={() => router.push("/login")}
                    className="w-full h-11 bg-gradient-to-r from-green-600 to-emerald-600 hover:from-green-700 hover:to-emerald-700 text-white font-semibold rounded-xl transition-all duration-200"
                  >
                    Voltar ao Login
                  </Button>

                  <Button
                    onClick={() => setIsSuccess(false)}
                    variant="outline"
                    className="w-full h-11 border-gray-300 text-gray-700 hover:bg-gray-50 rounded-xl"
                  >
                    Enviar Novamente
                  </Button>
                </div>
              </div>
            </div>
          </div>
        </div>
      </div>
    )
  }

  return (
    <div className="min-h-screen flex">
      {/* Left Side - Image (70%) */}
      <div className="hidden lg:flex lg:w-[70%] bg-gradient-to-br from-orange-600 via-red-600 to-pink-700 relative overflow-hidden">
        <div className="absolute inset-0 bg-black/20" />
        <div className="relative z-10 flex flex-col justify-center items-center p-12 text-white w-full">
          <div className="max-w-2xl text-center">
            <h1 className="text-5xl font-bold mb-8">Esqueceu sua senha?</h1>
            <p className="text-2xl opacity-90 mb-12 leading-relaxed">
              Não se preocupe! Digite seu email e enviaremos um link para redefinir sua senha
            </p>

            {/* Placeholder para imagem do undraw */}
            <div className="w-[500px] h-[400px] bg-white/10 rounded-3xl flex items-center justify-center backdrop-blur-sm mx-auto">
              <div className="text-center">
                <div className="w-40 h-40 bg-white/20 rounded-full mx-auto mb-6 flex items-center justify-center">
                  <Mail className="w-20 h-20 text-white/80" />
                </div>
              </div>
            </div>
          </div>
        </div>

        {/* Decorative elements */}
        <div className="absolute top-0 right-0 w-60 h-60 bg-white/10 rounded-full -translate-y-32 translate-x-32" />
        <div className="absolute bottom-0 left-0 w-48 h-48 bg-white/10 rounded-full translate-y-24 -translate-x-24" />
        <div className="absolute top-1/2 left-0 w-32 h-32 bg-white/5 rounded-full -translate-x-16" />
      </div>

      {/* Right Side - Form */}
      <div className="flex-1 lg:w-[30%] flex items-center justify-center p-6 bg-gray-50">
        <div className="w-full max-w-sm">
          {/* Header */}
          <div className="text-center mb-8">
            <div className="inline-flex items-center justify-center w-16 h-16 bg-gradient-to-r from-orange-600 to-red-600 rounded-2xl mb-6">
              <Mail className="w-8 h-8 text-white" />
            </div>
            <h2 className="text-2xl font-bold text-gray-900 mb-2">Esqueceu a Senha?</h2>
            <p className="text-gray-600 text-sm">Digite seu email para recuperar</p>
          </div>

          {/* Form */}
          <div className="bg-white rounded-2xl shadow-xl p-6 border border-gray-100">
            <form onSubmit={handleSubmit} className="space-y-5">
              {error && (
                <Alert variant="destructive" className="border-red-200 bg-red-50">
                  <AlertDescription className="text-red-800 text-sm">{error}</AlertDescription>
                </Alert>
              )}

              {/* Email Field */}
              <div className="space-y-2">
                <Label htmlFor="email" className="text-sm font-semibold text-gray-700">
                  Email
                </Label>
                <div className="relative">
                  <Mail className="absolute left-3 top-1/2 transform -translate-y-1/2 text-gray-400 w-4 h-4" />
                  <Input
                    id="email"
                    type="email"
                    placeholder="seu@email.com"
                    value={email}
                    onChange={(e) => setEmail(e.target.value)}
                    disabled={isLoading}
                    required
                    className="pl-10 h-11 border-gray-200 focus:border-orange-500 focus:ring-orange-500 rounded-xl text-sm"
                  />
                </div>
              </div>

              {/* Submit Button */}
              <Button
                type="submit"
                disabled={isLoading}
                className="w-full h-11 bg-gradient-to-r from-orange-600 to-red-600 hover:from-orange-700 hover:to-red-700 text-white font-semibold rounded-xl transition-all duration-200 transform hover:scale-[1.02] disabled:transform-none text-sm"
              >
                {isLoading ? (
                  <>
                    <Loader2 className="mr-2 h-4 w-4 animate-spin" />
                    Enviando...
                  </>
                ) : (
                  "Enviar Link de Recuperação"
                )}
              </Button>
            </form>

            {/* Info */}
            <div className="mt-6 p-3 bg-blue-50 rounded-xl border border-blue-200">
              <p className="text-xs text-blue-800">
                <strong>Dica:</strong> Verifique também sua pasta de spam caso não encontre o email na caixa de entrada.
              </p>
            </div>
          </div>

          {/* Back to Login */}
          <div className="text-center mt-6">
            <Button
              onClick={() => router.push("/login")}
              variant="ghost"
              className="text-gray-600 hover:text-gray-800 font-medium text-sm"
            >
              <ArrowLeft className="w-4 h-4 mr-2" />
              Voltar ao Login
            </Button>
          </div>
        </div>
      </div>
    </div>
  )
}
