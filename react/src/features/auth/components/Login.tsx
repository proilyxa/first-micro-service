import { FC } from "react";
import styles from "../../../components/ui/input/auth.module.css";
import { Link } from "react-router-dom";

const Login: FC = () => {
  return (
    <div className={"flex flex-col gap-4"}>
      <div className={"flex flex-col items-center gap-2"}>
        <h2 className={"text-4xl "}>Login</h2>
        <Link to={"/register"} className={"text-xs text-blue-700"}>
          no account?
        </Link>
      </div>

      <form className={""} action="">
        <div className={styles.formGroup}>
          <label htmlFor="email">Email</label>
          <input className={styles.input} id={"email"} type="text" />
        </div>

        <div className={styles.formGroup}>
          <label htmlFor="password">Password</label>
          <input className={styles.input} id={"password"} type="password" />
        </div>

        <button className={`mt-2 ${styles.btn}`}>Login</button>
      </form>
    </div>
  );
};

export default Login;
