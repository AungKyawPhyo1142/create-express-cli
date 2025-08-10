export const ENV = {
    APP: process.env.APP || '',
    CORS_ORIGIN: process.env.CORS_ORIGIN || true,
    DATABASE_URL: process.env.DATABASE_URL,
    JWT_SECRET: process.env.JWT_SECRET || '',
    NODE_ENV: (process.env.NODE_ENV || 'dev') as 'local' | 'dev' | 'production',
    PORT: process.env.PORT || 3000,
};

if (ENV.NODE_ENV === 'local' && !process.env.CORS_ORIGIN) {
    ENV.CORS_ORIGIN = 'http://localhost:8080';
}

if (ENV.NODE_ENV === 'dev' && process.env.DEV_DATABASE_URL) {
    ENV.DATABASE_URL = process.env.DEV_DATABASE_URL;
}