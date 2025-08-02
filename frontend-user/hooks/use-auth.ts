"use client"

import { useEffect, useState } from "react"
import { useRouter } from "next/navigation"
import { useAuthStore } from "@/lib/auth-store"

export function useAuth(requireAuth = true) {
  const router = useRouter()
  const { isAuthenticated, checkAuth, user, token, isLoading } = useAuthStore()
  const [isClient, setIsClient] = useState(false)

  useEffect(() => {
    setIsClient(true)
  }, [])

  useEffect(() => {
    if (!isClient) return

    const isValid = checkAuth()

    if (requireAuth && !isValid) {
      router.push("/login")
    } else if (!requireAuth && isValid) {
      router.push("/")
    }
  }, [requireAuth, checkAuth, router, isClient])

  return {
    user,
    token,
    isAuthenticated,
    isLoading,
  }
}
