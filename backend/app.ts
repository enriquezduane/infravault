import express, { Request, Response } from 'express';

const app = express();
const port: number = 3000;

// Keep backend routes simple - no /api prefix needed
app.get('/', (req: Request, res: Response) => {
    res.send('wassup nigger');
});

app.listen(port, () => {
    console.log(`App listening on port ${port}`);
});

