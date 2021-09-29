import React, { FunctionComponent } from "preact";
import { useState, useEffect } from "preact/hooks";
import color from "./color";
import logo from "../images/logo.png";
import { Global, css } from "@emotion/react";
import {
  Button,
  ButtonGroup,
  Card,
  SvgIcon,
  Tab,
  Tabs,
  TextField,
} from "@material-ui/core";
import { makeStyles } from "@material-ui/styles";
import styled from "@emotion/styled";
import font from "../font/ScheherazadeNewBold.woff2";
import LoginType from "../constant/loginType";
import loginText from "../constant/loginText";
import loginMethod from "../api/rest";

const isWidthLimited = document.body.offsetWidth > 400;

const useStyles = makeStyles(() => ({
  root: {
    background: color.buttonColor,
    '&:hover': {
      background: color.buttonColor
    },
    color: 'white',
    border: 0,
    borderRadius: 3,
    height: 48,
    padding: "0 30px",
  },
  default: {
    color: color.defaultSvgColor,
    backgroundColor: "transparent",
    minHeight: '30px',
    padding: isWidthLimited ? '0 10px' : '0 5px',
    minWidth: isWidthLimited ? '50px' : 'auto',
    marginRight: isWidthLimited  ? '20px' : 0,
    fontSize: isWidthLimited ? 'auto' : '4vmin'
  },
  tabs: {
    minHeight: '30px',
  }
}));

const Login: FunctionComponent<unknown> = () => {
  return (
    <Main id="Container">
      <Global
        styles={css`
          body {
            margin: 0;
            font-family: Gilroy-Medium, Tahoma, PingFangSC-Regular,
              Microsoft Yahei, Myriad Pro, Hiragino Sans GB, sans-serif;
          }
          @font-face {
            font-family: "Flower Font";
            src: url("${font}") format("woff2");
          }
        `}
      />
      <LoginForm />
    </Main>
  );
};

const Main = styled.div`
  height: ${window.innerHeight}px;
  background-image: ${color.backgroundColor};
  position: relative;
  display: flex;
  align-items: center;
  justify-content: center;
`;

const Logo = styled.img`
  width: 10vw;
  min-width: 150px;
  margin-bottom: 2rem;
`;

const LoginForm: FunctionComponent<unknown> = () => {
  const [loginType, setLoginType] = useState(LoginType.Phone);
  const tabIndex = {
    [LoginType.Phone]: 0,
    [LoginType.Email]: 1,
    [LoginType.SMS]: 2,
    [LoginType.LarkOauth]: 3,
  };
  const classes = useStyles();
  const loginFieldText = loginText[loginType];
  const loginForm = loginFieldText.map((item) => (
    <LoginInputArea
      id={item}
      key={item}
      label={item}
      variant={"outlined"}
    ></LoginInputArea>
  ));

  const login = () => {
    const opts: { [key: string]: string } = {};
    for (const id of loginFieldText) {
      const ele = document.querySelector(`#${id}`) as HTMLInputElement;
      const value = ele?.value;
      opts[id] = value;
    }
    loginMethod[loginType](opts, "null");
  };

  return (
    <FormLayout>
      <H2>登录Unique Studio</H2>
      {/* <Logo src={logo} /> */}
      <Tabs className={classes.tabs} value={tabIndex[loginType]}>
        <Tab
          label="手机号"
          onClick={() => setLoginType(LoginType.Phone)}
          className={classes.default}
        />
        <Tab
          label="邮箱"
          onClick={() => setLoginType(LoginType.Email)}
          className={classes.default}
        />
        <Tab
          label="短信验证码"
          onClick={() => setLoginType(LoginType.SMS)}
          className={classes.default}
        />
        <Tab
          label="Lark认证"
          onClick={() => setLoginType(LoginType.LarkOauth)}
          className={classes.default}
        />
      </Tabs>
      {loginForm}
      <LoginSubmitButton
        onClick={() => login()}
        className={classes.root}
        variant="text"
        disableFocusRipple
      >
        Next
      </LoginSubmitButton >
      <LoginSwitch></LoginSwitch>
    </FormLayout>
  );
};

const H2 = styled.h2`
  font-weight: 600;
  font-size: ${isWidthLimited ? 'auto' : '4vmin'};
`;

const FormLayout = styled(Card)`
  padding: 0 30px 30px 30px;
  width: 400px;
  display: flex;
  flex-direction: column;
  justify-content: center;
  height: auto;
  transition: height 0.5s;
`;

const LoginInputArea = styled(TextField)`
  margin: 0.5rem 0;
  max-width: 400px;
`;

const LoginSubmitButton = styled(Button)`
  margin: 2rem 0 1rem ;
`;

const LoginSwitch = styled.div`
  display: flex;
  align-items: center;
  justify-content: center;
  width: 100%;
  position: relative;
  max-width: 300px;
`;

export default Login;
