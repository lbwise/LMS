import express from 'express';
import expressSvelte from 'express-svelte'
const app = express();

app.set('view engine', 'html')
app.use(expressSvelte())

app.get('/', (req, res) => {
    res.svelte('./index')
});

app.listen(8080, () => {
    console.log('LISTENING ON PORT 8080')
});