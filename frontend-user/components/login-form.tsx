"use client"

import type React from "react"

import { useState } from "react"
import { useRouter } from "next/navigation"
import { useAuthStore } from "@/lib/auth-store"
import { Button } from "@/components/ui/button"
import { Input } from "@/components/ui/input"
import { Label } from "@/components/ui/label"
import { Alert, AlertDescription } from "@/components/ui/alert"
import { Loader2, Eye, EyeOff, Mail, Lock } from "lucide-react"

export default function LoginForm() {
  const [email, setEmail] = useState("")
  const [password, setPassword] = useState("")
  const [error, setError] = useState("")
  const [showPassword, setShowPassword] = useState(false)

  const { login, isLoading } = useAuthStore()
  const router = useRouter()

  const handleSubmit = async (e: React.FormEvent) => {
    e.preventDefault()
    setError("")

    if (!email || !password) {
      setError("Por favor, preencha todos os campos")
      return
    }

    try {
      const success = await login(email, password)

      if (success) {
        router.push("/")
      } else {
        setError("Credenciais inválidas")
      }
    } catch {
      setError("Erro ao fazer login. Tente novamente.")
    }
  }

  return (
    <div className="min-h-screen flex">
      {/* Left Side - Image (70%) */}
      <div className="hidden lg:flex lg:w-[70%] bg-gradient-to-br from-blue-600 via-purple-600 to-indigo-700 relative overflow-hidden">
        <div className="absolute inset-0 bg-black/20" />
        <div className="relative z-10 flex flex-col justify-center items-center p-12 text-white w-full">
          <div className="max-w-2xl text-center">
            <h1 className="text-5xl font-bold mb-8">Bem-vindo de volta!</h1>
            <p className="text-2xl opacity-90 mb-12 leading-relaxed">Acesse sua conta e continue de onde parou</p>

            {/* Placeholder para imagem */}
            <div className="w-[500px] h-[400px] bg-white/10 rounded-3xl flex items-center justify-center backdrop-blur-sm mx-auto">
              <div className="text-center">
                <div className="w-40 h-40 bg-white/20 rounded-full mx-auto mb-6 flex items-center justify-center">
                  <Lock className="w-20 h-20 text-white/80" />
                </div>
              </div>
            </div>
          </div>
        </div>

        {/* Decorative elements - maiores */}
        <div className="absolute top-0 right-0 w-60 h-60 bg-white/10 rounded-full -translate-y-32 translate-x-32" />
        <div className="absolute bottom-0 left-0 w-48 h-48 bg-white/10 rounded-full translate-y-24 -translate-x-24" />
        <div className="absolute top-1/2 left-0 w-32 h-32 bg-white/5 rounded-full -translate-x-16" />
        <div className="absolute top-1/4 right-1/4 w-24 h-24 bg-white/5 rounded-full" />
      </div>

      {/* Right Side - Login Form (30%) */}
      <div className="flex-1 lg:w-[30%] flex items-center justify-center p-6 bg-gray-50">
        <div className="w-full max-w-sm">
          {/* Header */}
          <div className="text-center mb-8">
            <div className="inline-flex items-center justify-center w-16 h-16 bg-gradient-to-r from-blue-600 to-purple-600 rounded-2xl mb-6">
              <Lock className="w-8 h-8 text-white" />
            </div>
            <h2 className="text-2xl font-bold text-gray-900 mb-2">Fazer Login</h2>
            <p className="text-gray-600 text-sm">Entre com suas credenciais</p>
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
                    className="pl-10 h-11 border-gray-200 focus:border-blue-500 focus:ring-blue-500 rounded-xl text-sm"
                  />
                </div>
              </div>

              {/* Password Field */}
              <div className="space-y-2">
                <Label htmlFor="password" className="text-sm font-semibold text-gray-700">
                  Senha
                </Label>
                <div className="relative">
                  <Lock className="absolute left-3 top-1/2 transform -translate-y-1/2 text-gray-400 w-4 h-4" />
                  <Input
                    id="password"
                    type={showPassword ? "text" : "password"}
                    placeholder="Sua senha"
                    value={password}
                    onChange={(e) => setPassword(e.target.value)}
                    disabled={isLoading}
                    required
                    className="pl-10 pr-10 h-11 border-gray-200 focus:border-blue-500 focus:ring-blue-500 rounded-xl text-sm"
                  />
                  <button
                    type="button"
                    onClick={() => setShowPassword(!showPassword)}
                    className="absolute right-3 top-1/2 transform -translate-y-1/2 text-gray-400 hover:text-gray-600 transition-colors"
                  >
                    {showPassword ? <EyeOff className="w-4 h-4" /> : <Eye className="w-4 h-4" />}
                  </button>
                </div>
              </div>

              {/* Remember Me & Forgot Password */}
              <div className="flex items-center justify-between">
                <label className="flex items-center">
                  <input
                    type="checkbox"
                    className="rounded border-gray-300 text-blue-600 focus:ring-blue-500 focus:ring-offset-0 w-4 h-4"
                  />
                  <span className="ml-2 text-xs text-gray-600">Lembrar</span>
                </label>
                <button
                  type="button"
				  onClick={() => router.push("/forgot-password")}
                  className="text-xs text-blue-600 hover:text-blue-700 font-medium transition-colors"
                >
                  Esqueceu?
                </button>
              </div>

              {/* Submit Button */}
              <Button
                type="submit"
                disabled={isLoading}
                className="w-full h-11 bg-gradient-to-r from-blue-600 to-purple-600 hover:from-blue-700 hover:to-purple-700 text-white font-semibold rounded-xl transition-all duration-200 transform hover:scale-[1.02] disabled:transform-none text-sm"
              >
                {isLoading ? (
                  <>
                    <Loader2 className="mr-2 h-4 w-4 animate-spin" />
                    Entrando...
                  </>
                ) : (
                  "Entrar"
                )}
              </Button>
            </form>


          </div>


        </div>
      </div>
    </div>
  )
}
