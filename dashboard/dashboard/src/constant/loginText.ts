import loginType from "./loginType";

export default {
  [loginType.Email]: ["Address", "Password", "oauth"],
  [loginType.Phone]: ["Number", "Password"],
  [loginType.SMS]: ["Number", "AuthCode"],
  [loginType.LarkOauth]: ["Unknown", "Unknown"],
};
