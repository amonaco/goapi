/* tslint:disable */
// api v0.0.1 b9d6b31ce576a5189ebbbc536282808a40261f3c
// --
// This file has been generated by https://github.com/webrpc/webrpc using gen/typescript
// Do not edit by hand. Update your webrpc schema and re-generate.

// WebRPC description and code-gen version
export const WebRPCVersion = "v1"

// Schema version of your RIDL schema
export const WebRPCSchemaVersion = "v0.0.1"

// Schema hash generated from your RIDL schema
export const WebRPCSchemaHash = "b9d6b31ce576a5189ebbbc536282808a40261f3c"


//
// Types
//
export interface Status {
  success: boolean
  error: string
  id: number
}

export interface Credentials {
  email: string
  password: string
  roles: Array<string>
  company_id: number
}

export interface PasswordToken {
  password: string
  token: string
}

export interface AuthService {
  signup(args: SignupArgs, headers?: object): Promise<SignupReturn>
  login(args: LoginArgs, headers?: object): Promise<LoginReturn>
  logout(args: LogoutArgs, headers?: object): Promise<LogoutReturn>
  setPassword(args: SetPasswordArgs, headers?: object): Promise<SetPasswordReturn>
  forgotPassword(args: ForgotPasswordArgs, headers?: object): Promise<ForgotPasswordReturn>
}

export interface SignupArgs {
  creds: Credentials
  signupToken?: string
}

export interface SignupReturn {
  authToken: string  
}
export interface LoginArgs {
  creds: Credentials
}

export interface LoginReturn {
  authToken: string  
}
export interface LogoutArgs {
  authToken: string
}

export interface LogoutReturn {
  status: Status  
}
export interface SetPasswordArgs {
  creds: PasswordToken
}

export interface SetPasswordReturn {
  status: Status  
}
export interface ForgotPasswordArgs {
  email: string
}

export interface ForgotPasswordReturn {  
}


  
//
// Client
//
export class AuthService implements AuthService {
  private hostname: string
  private fetch: Fetch
  private path = '/rpc/AuthService/'

  constructor(hostname: string, fetch: Fetch) {
    this.hostname = hostname
    this.fetch = fetch
  }

  private url(name: string): string {
    return this.hostname + this.path + name
  }
  
  signup = (args: SignupArgs, headers?: object): Promise<SignupReturn> => {
    return this.fetch(
      this.url('Signup'),
      createHTTPRequest(args, headers)).then((res) => {
      return buildResponse(res).then(_data => {
        return {
          authToken: <string>(_data.authToken)
        }
      })
    })
  }
  
  login = (args: LoginArgs, headers?: object): Promise<LoginReturn> => {
    return this.fetch(
      this.url('Login'),
      createHTTPRequest(args, headers)).then((res) => {
      return buildResponse(res).then(_data => {
        return {
          authToken: <string>(_data.authToken)
        }
      })
    })
  }
  
  logout = (args: LogoutArgs, headers?: object): Promise<LogoutReturn> => {
    return this.fetch(
      this.url('Logout'),
      createHTTPRequest(args, headers)).then((res) => {
      return buildResponse(res).then(_data => {
        return {
          status: <Status>(_data.status)
        }
      })
    })
  }
  
  setPassword = (args: SetPasswordArgs, headers?: object): Promise<SetPasswordReturn> => {
    return this.fetch(
      this.url('SetPassword'),
      createHTTPRequest(args, headers)).then((res) => {
      return buildResponse(res).then(_data => {
        return {
          status: <Status>(_data.status)
        }
      })
    })
  }
  
  forgotPassword = (args: ForgotPasswordArgs, headers?: object): Promise<ForgotPasswordReturn> => {
    return this.fetch(
      this.url('ForgotPassword'),
      createHTTPRequest(args, headers)).then((res) => {
      return buildResponse(res).then(_data => {
        return {
        }
      })
    })
  }
  
}

  
export interface WebRPCError extends Error {
  code: string
  msg: string
	status: number
}

const createHTTPRequest = (body: object = {}, headers: object = {}): object => {
  return {
    method: 'POST',
    headers: { ...headers, 'Content-Type': 'application/json' },
    body: JSON.stringify(body || {})
  }
}

const buildResponse = (res: Response): Promise<any> => {
  return res.text().then(text => {
    let data
    try {
      data = JSON.parse(text)
    } catch(err) {
      throw { code: 'unknown', msg: `expecting JSON, got: ${text}`, status: res.status } as WebRPCError
    }
    if (!res.ok) {
      throw data // webrpc error response
    }
    return data
  })
}

export type Fetch = (input: RequestInfo, init?: RequestInit) => Promise<Response>
