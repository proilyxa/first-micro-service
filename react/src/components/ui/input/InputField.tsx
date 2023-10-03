import React, { forwardRef } from "react";
import { InputType } from "@/components/ui/input/inputType.ts";
import styles from "@/components/ui/input/auth.module.css";
import cn from "classnames";

const InputField = forwardRef<HTMLInputElement, InputType>(
  (
    { type, errors = [], label, className, ...props }: InputType,
    ref: React.ForwardedRef<HTMLInputElement>,
  ) => {
    return (
      <div className={cn(styles.formGroup, className)}>
        <label htmlFor="firstName">{label}</label>
        <input
          className={styles.input}
          id="firstName"
          type={type}
          ref={ref}
          {...props}
        />
        {errors?.map((error: string, idx: number) => {
          return (
            <small className={"text-red-500"} key={idx}>
              {error}
            </small>
          );
        })}
      </div>
    );
  },
);

export default InputField;
