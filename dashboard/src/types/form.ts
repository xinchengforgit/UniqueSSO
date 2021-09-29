export interface EmailForm {
  Email: string;
  Password: string;
  [key: string]: string;
}

export interface PhoneForm {
  Phone: string;
  Password: string;
  [key: string]: string;
}

export interface SMSForm {
  Phone: string;
  Code: string;
  [key: string]: string;
}

export interface OauthForm {
  QRCodeSrc: string;
  [key: string]: string;
}

export interface LoginResponse {
  serviceResponse: {
    authenticationSuccess?: {
      redirectService: string;
    };
    authenticationFailure?: {
      description: string;
    };
  };
}

export type formType = EmailForm | SMSForm | OauthForm | LoginResponse;