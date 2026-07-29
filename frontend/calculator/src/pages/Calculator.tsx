import React, { useState } from 'react'
import { calculate } from '../api/calculator'
import TraditionalCalculator from './TraditionalCalculator'
import ModernCalculator from './ModernCalculator'
import '../App.css'

export default function Calculator() {
  const [active, setActive] = useState<'both' | 'traditional' | 'modern'>('both')
  const [mobileView, setMobileView] = useState<'left' | 'right'>('left')

  return (
    <div>
      <h1 className="calc-title">Calculator — Two Styles</h1>
      <div className="dual-calculator-root">


        <div className="dual-wrap">
          <section className={`pane left ${active === 'modern' ? 'dim' : ''} ${mobileView === 'left' ? 'mobile-visible' : ''}`}>
            <TraditionalCalculator />
          </section>

          <section className={`pane right ${active === 'traditional' ? 'dim' : ''} ${mobileView === 'right' ? 'mobile-visible' : ''}`}>
            <ModernCalculator />
          </section>
        </div>
        <div className="mobile-controls">
          <button onClick={() => setMobileView('left')} aria-pressed={mobileView === 'left'}>Traditional</button>
          <button onClick={() => setMobileView('right')} aria-pressed={mobileView === 'right'}>Modern</button>
        </div>

        <div style={{ position: 'absolute', left: '-9999px', top: '-9999px' }}>
          <LegacyForm />
        </div>
      </div>
    </div>
  )
}

function LegacyForm() {
  const [a, setA] = useState('')
  const [b, setB] = useState('')
  const [op, setOp] = useState('add')
  const [loading, setLoading] = useState(false)
  const [result, setResult] = useState<number | null>(null)
  const [error, setError] = useState<string | null>(null)

  const operations = [
    { value: 'add', label: 'Addition (+)' },
    { value: 'subtract', label: 'Subtraction (-)' },
    { value: 'multiply', label: 'Multiplication (×)' },
    { value: 'divide', label: 'Division (÷)' },
    { value: 'exponent', label: 'Exponentiation (^)' },
    { value: 'sqrt', label: 'Square Root (√)' },
    { value: 'percentage', label: 'Percentage (%)' },
  ]

  const handleCalculate = async () => {
    setError(null)
    setResult(null)

    if (a.trim() === '') {
      setError('Value A is required')
      return
    }

    const aNum = Number(a)
    if (Number.isNaN(aNum)) {
      setError('Value A must be a valid number')
      return
    }

    let bNum = 0
    if (op !== 'sqrt') {
      if (b.trim() === '') {
        setError('Value B is required for this operation')
        return
      }
      bNum = Number(b)
      if (Number.isNaN(bNum)) {
        setError('Value B must be a valid number')
        return
      }
    }

    if (op === 'divide' && bNum === 0) {
      setError('Division by zero is not allowed')
      return
    }

    if (op === 'sqrt' && aNum < 0) {
      setError('Square root of negative number is not allowed')
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
    <div className="legacy-form" style={{ marginTop: 20 }}>
      <div className="form-row">
        <label htmlFor="value-a">Value A</label>
        <input
          id="value-a"
          aria-label="value-a"
          type="text"
          value={a}
          onChange={(e) => setA(e.target.value)}
        />
      </div>

      <div className="form-row">
        <label htmlFor="value-b">Value B</label>
        <input
          id="value-b"
          aria-label="value-b"
          type="text"
          value={b}
          onChange={(e) => setB(e.target.value)}
          disabled={op === 'sqrt'}
          placeholder={op === 'sqrt' ? 'Not required for square root' : ''}
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
