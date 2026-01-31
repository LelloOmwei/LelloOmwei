use serde::{Serialize};
use std::time::{SystemTime, UNIX_EPOCH};

#[derive(Serialize)]
struct Token {
    id: String,
    sig: String,
    val: f64,
    ts: u64,
}

#[tokio::main]
async fn main() -> Result<(), Box<dyn std::error::Error>> {
    let client = reqwest::Client::new();
    
    // Získanie aktuálneho času
    let start = SystemTime::now();
    let since_the_epoch = start.duration_since(UNIX_EPOCH)?;
    
    // Vytvorenie tvojho sémantického tokenu
    let token = Token {
        id: "rust-edge-01".to_string(),
        sig: "konesuhladne".to_string(),
        val: 1.0,
        ts: since_the_epoch.as_secs(),
    };

    println!("🚀 Posielam token z Rustu: {} -> {}", token.id, token.sig);

    let res = client.post("http://localhost:8080/ingest")
        .json(&token)
        .send()
        .await?;

    if res.status().is_success() {
        println!("✅ Token prijatý bránou!");
    } else {
        println!("❌ Chyba pri doručení: {:?}", res.status());
    }

    Ok(())
}
