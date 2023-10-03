import { InputHTMLAttributes } from "react";

export interface InputType extends InputHTMLAttributes<HTMLInputElement> {
  errors?: string[] | null;
  label: string;
}
