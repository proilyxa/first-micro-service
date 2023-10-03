import { JSX } from "react";
import { Outlet } from "react-router-dom";

export default function Home(): JSX.Element {
  return (
    <div className={"w-full"}>
      <Outlet />
    </div>
  );
}
