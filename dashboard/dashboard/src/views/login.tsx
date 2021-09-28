import React, { FunctionComponent } from 'preact'
import {useState} from 'preact/hooks'
import color from './color'
import logo from '../images/logo.png'
import { Global, css } from '@emotion/react';
import { Button, Card, SvgIcon, TextField } from '@material-ui/core';
import { makeStyles } from '@material-ui/styles'
import styled from '@emotion/styled';
import mail from '@material-ui/icons/Mail'
import phone from '@material-ui/icons/Phone'
import sms from '@material-ui/icons/Sms'
import oauth from '@material-ui/icons/Vpnkey'
import font from '../font/DutchTulips.woff2'
import loginType from '../constant/loginType'
import loginText from '../constant/loginText'

interface AppProps {}

const useStyles = makeStyles((theme ?: any) => ({
  root: {
    background: color.buttonColor,
    border: 0,
    borderRadius: 3,
    color: 'white',
    height: 48,
    padding: '0 30px',
  },
  default: {
    color: color.defaultSvgColor,
    borderRadius: '100%',
    height: 64,
    width: 64
  }
}));


const Login : FunctionComponent<AppProps> = () => {
  return <Main id='Container'>
    <Global styles={css`
      body{
        margin: 0;
      }
      @font-face {
        font-family: "Flower Font";
        src: url("${font}") format("woff2");
      }
    `}/>
    <Logo src={logo} />
    <LoginForm></LoginForm>
  </Main>
}

const Main = styled.div`
  height: 100vh;
  background-image: ${color.backgroundColor};
  position: relative;
  display: flex;
  align-items: center;
  justify-content: center;
`

const Logo = styled.img`
  position: absolute;
  top: 3vw;
  left: 3vw;
  width: 12vw;
  min-width: 150px;
`

const LoginForm : FunctionComponent<AppProps> = () => {
  const [login, setLogin] = useState(loginType.Phone);
  const classes = useStyles()
  const text = loginText[login];
  const loginForm = text.map( item => <LoginInputArea label={item} variant={'standard'}></LoginInputArea>)
  return <FormLayout>
    <LoginText>Login</LoginText>
    {loginForm}
    <LoginSubmitButton className={classes.root} variant="text">Submit!</LoginSubmitButton>
    <LoginSwitch>
    <Button onClick={() => setLogin(loginType.Phone)} className={classes.default} variant="text"><SvgIcon component={phone}></SvgIcon></Button>
    <Button onClick={() => setLogin(loginType.Email)} className={classes.default} variant="text"><SvgIcon component={mail}></SvgIcon></Button>
    <Button onClick={() => setLogin(loginType.SMS)} className={classes.default} variant="text"><SvgIcon component={sms}></SvgIcon></Button>
    <Button onClick={() => setLogin(loginType.LarkOauth)} className={classes.default} variant="text"><SvgIcon component={oauth}></SvgIcon></Button>
    </LoginSwitch>
  </FormLayout>
}

const FormLayout = styled(Card)`
  padding: 50px;
  width: 500px;
  display: flex;
  flex-direction: column;
  align-items: center;
  justify-content: center;
  transition: height 0.5s ;
`

const LoginText = styled.div`
  color: ${color.loginFontColor};
  font-size: 2rem;
  letter-spacing: calc(1rem + 5px);
  text-align: center;
  text-indent: calc(1rem + 5px);
  font-style: italic;
  font-family: 'Flower Font';
  margin-bottom: 1rem;
`

const LoginInputArea = styled(TextField)`
  margin: 0.5rem;
`

const LoginSubmitButton = styled(Button)`
  margin: 2rem 2rem 1rem 2rem;
`

const LoginSwitch = styled.div`
  display: flex;
  align-items: center;
  justify-content: space-between;
  width: 60%;
`

export default Login;