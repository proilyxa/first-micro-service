import { createBrowserRouter } from "react-router-dom";
import Login from "@/features/auth/components/Login.tsx";
import Register from "@/features/auth/components/Register.tsx";
import AuthLayout from "@/pages/Auth.tsx";
import Home from "@/pages/Home.tsx";
import Dashboard from "@/features/home/components/Dashboard.tsx";
import Articles from "@/features/articles/components/Articles.tsx";

const router = createBrowserRouter([
  {
    path: "/",
    element: <Home />,
    children: [
      {
        path: "/",
        element: <Dashboard />,
      },
      {
        path: "/articles",
        element: <Articles />,
      },
    ],
  },
  {
    element: <AuthLayout />,
    children: [
      {
        path: "login",
        element: <Login />,
      },
      {
        path: "register",
        element: <Register />,
      },
    ],
  },
]);

export default router;
