import { EmailForm, OauthForm, PhoneForm, SMSForm } from "../types/form";
import loginType from '../constant/loginType'
import conf from "../config/conf";

interface RestForm {
  [key: string]: string
}

const Endpoint = {
  emailLogin: (service: string) => `/cas/login?type=email&service=${service}`,
  phoneLogin: (service: string) => `/cas/login?type=phone&service=${service}`,
  smsLogin: (service: string) => `/cas/login?type=sms&service=${service}`,
  larkLogin: (service: string) => `/cas/login?type=lark&service=${service}`,
};

const wrapLoginRequest = async (
  url: string,
  jsonData: Record<string, unknown>
): Promise<string> => {
  //initialize fetch options
  const opts: RequestInit = {
    method: "POST",
    mode: "cors",
    headers: {
      "Content-Type": "application/json",
    },
    body: JSON.stringify(jsonData),
  };
  try {
    const data = await fetch(`${conf.baseUrl}${url}`, opts).then((result) =>
      result.json()
    );
    if (data.serviceResponse.authenticationSuccess !== undefined) {
      return data.serviceResponse.authenticationSuccess.redirectService;
    } else {
      throw Error("response");
    }
  } catch ({ response }) {
    console.log(response);
    throw Error("send http request failed");
  }
};

export default {
  [loginType.Email]: async (form: RestForm, service: string): Promise<string> => {
    return wrapLoginRequest(Endpoint.emailLogin(service), {
      email: form.Email,
      password: form.Password,
    });
  },

  [loginType.Phone]: async (form: RestForm, service: string): Promise<string> => {
    return wrapLoginRequest(Endpoint.phoneLogin(service), {
      phone: form.Phone,
      password: form.Password,
    });
  },

  [loginType.SMS]: async (form: RestForm, service: string): Promise<string> => {
    return wrapLoginRequest(Endpoint.smsLogin(service), {
      phone: form.Phone,
      code: form.Code,
    });
  },

  [loginType.LarkOauth]: async (form: RestForm, service: string): Promise<string> => {
    return wrapLoginRequest(Endpoint.larkLogin(service), {
      qrcode_src: form.QRCodeSrc,
    });
  },
};
