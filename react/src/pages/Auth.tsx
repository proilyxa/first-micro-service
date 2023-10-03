import { JSX } from "react";
import { Outlet } from "react-router-dom";

export default function Auth(): JSX.Element {
  return (
    <div className={'flex justify-center h-screen pt-44'}>
        <div className={'w-3/12'}>
            <Outlet />
        </div>
    </div>
  );
}
