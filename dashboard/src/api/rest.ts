import LoginType from "@/constant/loginType";
import {
  EmailForm,
  LoginResponse,
  OauthForm,
  PhoneForm,
  SMSForm,
} from "@/model/form";

const Endpoint =  {
  emailLogin: (service: string) =>
    `/cas/login?type=${LoginType.Email}&service=${service}`,
  phoneLogin: (service: string) =>
    `/cas/login?type=${LoginType.Phone}&service=${service}`,
  smsLogin: (service: string) =>
    `/cas/login?type=${LoginType.SMS}&service=${service}`,
  larkLogin: (service: string) =>
    `/cas/login?type=${LoginType.LarkOauth}&service=${service}`,
}

export class RestClient {
  baseUrl : string;
  constructor(baseUrl: string) {
    this.baseUrl = baseUrl;
  }
  private wrapLoginRequest = async(
    url: string,
    jsonData: Record<string, unknown>
  ): Promise<string> => {
    //initialize fetch options
    const opts : RequestInit = {
      method: 'POST',
      mode: 'cors',
      headers: {
        'Content-Type': 'application/json'
      },
      body: JSON.stringify(jsonData)
    }
    try {
      const data = await fetch(`${this.baseUrl}${url}`, opts).then(result => result.json())
      if (data.serviceResponse.authenticationSuccess !== undefined) {
        return data.serviceResponse.authenticationSuccess.redirectService;
      } else {
        throw Error("response");
      }
    } catch ({ response }) {
      console.log(response);
      throw Error("send http request failed");
    }
  }

  loginByEmail = async(form: EmailForm, service: string): Promise<string> => {
    return this.wrapLoginRequest(Endpoint.emailLogin(service), {
      email: form.Email,
      password: form.Password,
    });
  }

  loginByPhone = async(form: PhoneForm, service: string): Promise<string> => {
    return this.wrapLoginRequest(Endpoint.phoneLogin(service), {
      phone: form.Phone,
      password: form.Password,
    });
  }

  loginBySMS = async(form: SMSForm, service: string): Promise<string> => {
    return this.wrapLoginRequest(Endpoint.smsLogin(service), {
      phone: form.Phone,
      code: form.Code,
    });
  }

  LoginByOauth = async(form: OauthForm, service: string): Promise<string> => {
    return this.wrapLoginRequest(Endpoint.larkLogin(service), {
      qrcode_src: form.QRCodeSrc,
    });
  }
}
