import Head from 'next/head'
import Image from 'next/image'
import styles from '../styles/Home.module.css'

export default function Home() {
  return (
    <div className={styles.container}>
      <Head>
        <title>Kioku</title>
        <meta name="description" content="Kioku" />
        <link rel="icon" href="/favicon.ico" />
      </Head>

      <main className={styles.main}>
        <h1 className={styles.title}>
          Hello, World 👋
        </h1>
      </main>

      <footer className={styles.footer}>
        <a
          href="https://github.com/kioku-project/kioku"
          target="_blank"
          rel="noopener noreferrer"
        >
        <Image src="/github.png" alt="GitHub Icon" width={16} height={16} />
        </a>
      </footer>
    </div>
  )
}
