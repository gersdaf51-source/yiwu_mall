export default async function Home() {
  const res = await fetch("http://localhost:8080/ping")
  const data = await res.json()

  return (
    <div>
      <h1>Yiwu Mall</h1>
      <pre>{JSON.stringify(data, null, 2)}</pre>
    </div>
  )
}