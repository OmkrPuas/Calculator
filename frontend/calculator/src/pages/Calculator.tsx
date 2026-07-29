import { useState } from 'react'
import { calculate } from '../api/calculator'
import '../App.css'

const operations = [
  { value: 'add', label: 'Addition (+)' },
  { value: 'subtract', label: 'Subtraction (-)' },
  { value: 'multiply', label: 'Multiplication (×)' },
  { value: 'divide', label: 'Division (÷)' },
]

export default function Calculator() {
  const [a, setA] = useState('')
  const [b, setB] = useState('')
  const [op, setOp] = useState('add')
  const [loading, setLoading] = useState(false)
  const [result, setResult] = useState<number | null>(null)
  const [error, setError] = useState<string | null>(null)

  const handleCalculate = async () => {
    setError(null)
    setResult(null)

    if (a.trim() === '' || b.trim() === '') {
      setError('Both inputs are required')
      return
    }

    const aNum = Number(a)
    const bNum = Number(b)
    if (Number.isNaN(aNum) || Number.isNaN(bNum)) {
      setError('Inputs must be valid numbers')
      return
    }

    if (op === 'divide' && bNum === 0) {
      setError('Division by zero is not allowed')
      return
    }

    setLoading(true)
    try {
      const res = await calculate({ operation: op, a: aNum, b: bNum })
      setResult(res)
    } catch (err: any) {
      setError(err?.message || 'Unexpected error')
    } finally {
      setLoading(false)
    }
  }

  return (
    <div className="calculator-root">
      <h1>Calculator</h1>

      <div className="form-row">
        <label>Value A</label>
        <input
          aria-label="value-a"
          type="text"
          value={a}
          onChange={(e) => setA(e.target.value)}
        />
      </div>

      <div className="form-row">
        <label>Value B</label>
        <input
          aria-label="value-b"
          type="text"
          value={b}
          onChange={(e) => setB(e.target.value)}
        />
      </div>

      <div className="form-row">
        <label>Operation</label>
        <select value={op} onChange={(e) => setOp(e.target.value)}>
          {operations.map((o) => (
            <option key={o.value} value={o.value}>
              {o.label}
            </option>
          ))}
        </select>
      </div>

      <div className="form-row">
        <button type="button" onClick={handleCalculate} disabled={loading}>
          {loading ? 'Calculating…' : 'Calculate'}
        </button>
      </div>

      <div className="result-row">
        {error && <div className="error">Error: {error}</div>}
        {result !== null && <div className="result">Result: {result}</div>}
      </div>
    </div>
  )
}
