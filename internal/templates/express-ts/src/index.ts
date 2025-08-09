import express from "express";
import exampleRouter from "./routes/example";

const app = express();

app.use(express.json());

app.use("/api", exampleRouter);

const PORT = process.env.PORT || 3000;
app.listen(PORT, () => {
    console.log(`🚀 Server is running on http://localhost:${PORT}`);
});
