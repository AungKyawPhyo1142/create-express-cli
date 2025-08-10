export const ErrorCodes = {
  AUTHENTICATION_ERROR: {
    code: '{{ .ProjectName }}-40102',
    message: 'Authentication failed.',
    statusCode: 401,
    userMessage: 'Authentication failed. Please check your credentials.',
  },
  BAD_REQUEST: {
    code: '{{ .ProjectName }}-40001',
    message: 'Bad request.',
    statusCode: 400,
    userMessage: 'The request could not be understood by the server.',
  },
  CONFLICT: {
    code: '{{ .ProjectName }}-40901',
    message: 'Conflict occurred.',
    statusCode: 409,
    userMessage: 'There is a conflict with the current state of the resource.',
  },
  DATABASE_ERROR: {
    code: '{{ .ProjectName }}-50002',
    message: 'Database error.',
    statusCode: 500,
    userMessage: 'A database error occurred. Please try again later.',
  },
  EMAIL_VERIFICATION_ERROR: {
    code: '{{ .ProjectName }}-50003',
    message: 'Email verification error.',
    statusCode: 500,
    userMessage:
      'There was an error with email verification. Please try again later.',
  },
  FORBIDDEN: {
    code: '{{ .ProjectName }}-40301',
    message: 'Forbidden resource.',
    statusCode: 403,
    userMessage: 'You do not have permission to access this resource.',
  },
  INTERNAL_SERVER_ERROR: {
    code: '{{ .ProjectName }}-50001',
    message: 'Internal server error.',
    statusCode: 500,
    userMessage: 'An unexpected error occurred. Please try again later.',
  },
  NOT_FOUND: {
    code: '{{ .ProjectName }}-40401',
    message: 'Resource not found.',
    statusCode: 404,
    userMessage: 'The requested resource could not be found.',
  },
  SERVICE_UNAVAILABLE: {
    code: '{{ .ProjectName }}-50301',
    message: 'Service unavailable.',
    statusCode: 503,
    userMessage:
      'The service is currently unavailable. Please try again later.',
  },
  UNAUTHORIZED: {
    code: '{{ .ProjectName }}-40101',
    message: 'Unauthorized access.',
    statusCode: 401,
    userMessage: 'You are not authorized to perform this action.',
  },
  VALIDATION_ERROR: {
    code: '{{ .ProjectName }}-40002',
    message: 'Validation failed.',
    statusCode: 400,
    userMessage: 'Some of the provided data is invalid.',
  },
};
